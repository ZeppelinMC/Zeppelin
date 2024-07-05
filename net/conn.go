package net

import (
	"aether/net/io"
	"aether/net/packet"
	"aether/net/packet/handshake"
	"aether/net/packet/login"
	"aether/net/packet/status"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"unicode/utf16"
)

type Conn struct {
	net.Conn

	listener *Listener

	reader io.Reader
	writer io.Writer

	loginData login.LoginSuccess

	state int
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

		rd = io.NewReader(bytes.NewReader(data))
	} else {
		//compression
	}

	var pk packet.Packet
	pc, ok := serverboundPool[conn.state][packetId]
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
		return false
	}

	switch handshaking.NextState {
	case handshake.Status:
		conn.state = StatusState
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
		conn.state = LoginState
		pk, err := conn.ReadPacket()
		if err != nil {
			return false
		}
		loginStart, ok := pk.(*login.LoginStart)
		if !ok {
			return false
		}
		conn.loginData = login.LoginSuccess{
			UUID:     loginStart.PlayerUUID,
			Username: loginStart.Name,
		}
		if err := conn.WritePacket(&conn.loginData); err != nil {
			return false
		}
	}
	return true
}
