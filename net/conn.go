package net

import (
	"bufio"
	"bytes"
	"crypto/cipher"
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

	writer io.Writer

	username   string
	uuid       uuid.UUID
	properties []login.Property

	state atomic.Int32

	encrypted                 bool
	sharedSecret, verifyToken []byte

	rd  *bufio.Reader
	buf [4096]byte

	decrypter, encrypter cipher.Stream
}

func (conn *Conn) Username() string {
	return conn.username
}

func (conn *Conn) UUID() uuid.UUID {
	return conn.uuid
}

func (conn *Conn) Properties() []login.Property {
	return conn.properties
}

func (conn *Conn) SetState(state int32) {
	conn.state.Store(state)
}
func (conn *Conn) State() int32 {
	return conn.state.Load()
}

func (conn *Conn) WritePacket(pk packet.Packet) error {
	if conn.listener.CompressionThreshold < 0 {
		var packetBuf = new(bytes.Buffer)
		w := io.NewWriter(packetBuf)
		if err := w.VarInt(pk.ID()); err != nil {
			return err
		}
		if err := pk.Encode(w); err != nil {
			return err
		}
		data := packetBuf.Bytes()
		length := io.AppendVarInt(nil, int32(len(data)))

		_, err := conn.Write(append(length, data...))
		return err
	} else {
		//compressed
	}
	return nil
}

func (conn *Conn) Read(dst []byte) (i int, err error) {
	i, err = conn.Conn.Read(dst)
	if err != nil {
		return i, err
	}
	if conn.encrypted {
		conn.decrypt(dst, dst)
	}

	return i, err
}

func (conn *Conn) Write(data []byte) (i int, err error) {
	if !conn.encrypted {
		return conn.Conn.Write(data)
	}
	conn.encrypt(data, data)
	if err != nil {
		return 0, err
	}

	return conn.Conn.Write(data)
}

func (conn *Conn) ReadPacket() (packet.Packet, error) {
	var rd = io.NewReader(conn, 0)
	var (
		length, packetId int32
		data             []byte
	)
	if conn.listener.CompressionThreshold < 0 {
		if _, err := rd.VarInt(&length); err != nil {
			return nil, err
		}
		rd.SetLength(int(length))
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
		conn.username = loginStart.Name
		conn.uuid = loginStart.PlayerUUID

		if conn.listener.Encrypt {
			if err := conn.Encrypt(); err != nil {
				return false
			}
		}

		suc := &login.LoginSuccess{
			UUID:                conn.uuid,
			Username:            conn.username,
			Properties:          conn.properties,
			StrictErrorHandling: true,
		}
		if err := conn.WritePacket(suc); err != nil {
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
