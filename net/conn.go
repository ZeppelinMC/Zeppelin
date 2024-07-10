package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync/atomic"
	"unicode/utf16"

	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/handshake"
	"github.com/dynamitemc/aether/net/packet/login"
	"github.com/dynamitemc/aether/net/packet/status"

	"github.com/google/uuid"
)

type Conn struct {
	net.Conn

	listener *Listener

	reader io.Reader
	writer io.Writer

	loginData login.LoginSuccess

	state atomic.Int32
}

func (conn *Conn) Username() string {
	return conn.loginData.Username
}

func (conn *Conn) UUID() uuid.UUID {
	return conn.loginData.UUID
}

func (conn *Conn) Properties() []login.Property {
	return conn.loginData.Properties
}

func (conn *Conn) SetState(state int32) {
	conn.state.Store(state)
}
func (conn *Conn) State() int32 {
	return conn.state.Load()
}

func (conn *Conn) WritePacket(pk packet.Packet) error {
	if conn.listener.CompressionThreshold < 0 {
		var buf = new(bytes.Buffer)
		w := io.NewWriter(buf)
		if err := w.VarInt(pk.ID()); err != nil {
			return err
		}
		if err := pk.Encode(w); err != nil {
			return err
		}
		if err := conn.writer.VarInt(int32(buf.Len())); err != nil {
			return err
		}

		return conn.writer.FixedByteArray(buf.Bytes())
	} else {
		//compressed
	}
	return nil
}

func (conn *Conn) ReadPacket() (packet.Packet, error) {
	rd := conn.reader
	var (
		length, packetId int32
		data             []byte
	)
	if conn.listener.CompressionThreshold < 0 {
		if _, err := rd.VarInt(&length); err != nil {
			return nil, err
		}
		vii, err := rd.VarInt(&packetId)
		if err != nil {
			return nil, err
		}
		length -= int32(vii)
		data = make([]byte, length)
		if err := rd.FixedByteArray(data); err != nil {
			return nil, err
		}

		rd = io.NewReader(bytes.NewReader(data), int(length))
	} else {
		//compression
	}

	var pk packet.Packet
	pc, ok := serverboundPool[conn.state.Load()][packetId]
	if !ok {
		return packet.UnknownPacket{
			Id:      packetId,
			Length:  length,
			Payload: rd,
		}, nil
	} else {
		pk = pc()
		err := pk.Decode(rd)
		return pk, err
	}
}

func (conn *Conn) writeLegacyStatus(status status.StatusResponseData) {
	protocolString := fmt.Sprint(status.Version.Protocol)
	onlineString := fmt.Sprint(status.Players.Online)
	maxString := fmt.Sprint(status.Players.Max)

	stringData := make([]rune, 3, 3+len(protocolString)+len(status.Version.Name)+len(status.Description.Text)+len(onlineString)+len(maxString))
	stringData[0], stringData[1] = '§', '1'
	stringData = append(append(stringData, []rune(protocolString)...), 0)
	stringData = append(append(stringData, []rune(status.Version.Name)...), 0)
	stringData = append(append(stringData, []rune(status.Description.Text)...), 0)
	stringData = append(append(stringData, []rune(onlineString)...), 0)
	stringData = append(stringData, []rune(maxString)...)

	utf16be := utf16.Encode([]rune(stringData))

	length := uint16(len(utf16be))
	conn.Write([]byte{
		0xFF, byte(length >> 8), byte(length),
	})

	binary.Write(conn, binary.BigEndian, utf16be)
}

func (conn *Conn) writeLegacyDisconnect(reason string) {
	var data = []byte{0xFF}

	var strdata = utf16.Encode([]rune(reason))
	var strlen = int16(len(strdata))

	data = append(data, byte(strlen>>8), byte(strlen))
	conn.Write(data)
	binary.Write(conn, binary.BigEndian, strdata)
}

func (conn *Conn) handleHandshake() bool {
	pk, err := conn.ReadPacket()
	if err != nil {
		return false
	}
	handshaking, ok := pk.(*handshake.Handshaking)
	if !ok {
		if pk.ID() == 122 {
			conn.writeLegacyStatus(conn.listener.Status())
		}
		if pk.ID() == 78 {
			conn.writeLegacyDisconnect("Your client is too old! this server supports MC 1.21")
		}
		return false
	}

	switch handshaking.NextState {
	case handshake.Status:
		conn.state.Store(StatusState)
		pk, err := conn.ReadPacket()
		if err != nil {
			return false
		}
		_, ok := pk.(*status.StatusRequest)
		if !ok {
			return false
		}
		if err := conn.WritePacket(&status.StatusResponse{Data: conn.listener.Status()}); err != nil {
			return false
		}

		pk, err = conn.ReadPacket()
		if err != nil {
			return false
		}
		p, ok := pk.(*status.Ping)
		if !ok {
			return false
		}
		if err := conn.WritePacket(p); err != nil {
			return false
		}
	case handshake.Login:
		conn.state.Store(LoginState)
		pk, err := conn.ReadPacket()
		if err != nil {
			return false
		}
		loginStart, ok := pk.(*login.LoginStart)
		if !ok {
			return false
		}
		conn.loginData = login.LoginSuccess{
			UUID:                loginStart.PlayerUUID,
			Username:            loginStart.Name,
			StrictErrorHandling: true,
		}
		if err := conn.WritePacket(&conn.loginData); err != nil {
			return false
		}
		pk, err = conn.ReadPacket()
		if err != nil {
			return false
		}
		_, ok = pk.(*login.LoginAcknowledged)
		if !ok {
			return false
		}
		conn.state.Store(ConfigurationState)
		return true
	}
	return false
}
