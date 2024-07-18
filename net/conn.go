package net

import (
	"bytes"
	"compress/zlib"
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
	"github.com/dynamitemc/aether/text"

	"github.com/google/uuid"
)

const (
	clientVeryOldMsg = "Your client is WAYYYYYY too old!!! this server supports MC 1.21"
	clientTooOldMsg  = "Your client is too old! this server supports MC 1.21"
	clientTooNewMsg  = "Your client is too new! this server supports MC 1.21"
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

	decrypter, encrypter cipher.Stream
	compressionSet       bool
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
	var packetBuf = new(bytes.Buffer)
	w := io.NewWriter(packetBuf)
	if err := w.VarInt(pk.ID()); err != nil {
		return err
	}
	if err := pk.Encode(w); err != nil {
		return err
	}
	data := packetBuf.Bytes()

	if conn.listener.cfg.CompressionThreshold < 0 || !conn.compressionSet { // no encryption
		length := io.AppendVarInt(nil, int32(len(data)))
		_, err := conn.Write(append(length, data...))
		return err
	} else { // with encryption
		if len(data) < int(conn.listener.cfg.CompressionThreshold) {
			data = append([]byte{0}, data...)
			packetLength := io.AppendVarInt(nil, int32(len(data)))
			_, err := conn.Write(append(packetLength, data...))
			return err
		}

		dataLength := io.AppendVarInt(nil, int32(len(data)))

		var packetBuf = new(bytes.Buffer)
		comp := zlib.NewWriter(packetBuf)
		if _, err := comp.Write(data); err != nil {
			return err
		}
		if err := comp.Close(); err != nil {
			return err
		}

		compData := append(dataLength, packetBuf.Bytes()...)

		packetLength := io.AppendVarInt(nil, int32(len(compData)))

		_, err := conn.Write(append(packetLength, compData...))
		return err
	}
}

func (conn *Conn) Read(dst []byte) (i int, err error) {
	i, err = conn.Conn.Read(dst)
	if err != nil {
		return i, err
	}
	if conn.encrypted {
		conn.decryptd(dst, dst)
	}

	return i, err
}

func (conn *Conn) Write(data []byte) (i int, err error) {
	if !conn.encrypted {
		return conn.Conn.Write(data)
	}
	conn.encryptd(data, data)
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
	if conn.listener.cfg.CompressionThreshold < 0 || !conn.compressionSet {
		if _, err := rd.VarInt(&length); err != nil {
			return nil, err
		}
		vii, err := rd.VarInt(&packetId)
		if err != nil {
			return nil, err
		}
		length -= int32(vii)

		rd.SetLength(int(length))
		if length < 0 {
			return nil, fmt.Errorf("malformed packet, you are being fooled")
		}

		data = make([]byte, length)
		if err := rd.FixedByteArray(data); err != nil {
			return nil, err
		}

		rd = io.NewReader(bytes.NewReader(data), int(length))
	} else {
		var packetLength int32
		if _, err := rd.VarInt(&packetLength); err != nil {
			return nil, err
		}
		dli, err := rd.VarInt(&length)
		if err != nil {
			return nil, err
		}
		if length == 0 {
			length = packetLength - 1

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
			panic("implement compression reading")
			l := packetLength - int32(dli)
			var packetData = make([]byte, l)
			if _, err := conn.Read(packetData); err != nil {
				return nil, err
			}

			compressedData := bytes.NewReader(packetData)
			z, err := zlib.NewReader(compressedData)
			if err != nil {
				return nil, err
			}
			defer z.Close()
			var uncompressedData = make([]byte, length)
			fmt.Println(length, compressedData.Len())
			if _, err := z.Read(uncompressedData); err != nil {
				fmt.Println(err)
				return nil, err
			}

			rd = io.NewReader(bytes.NewReader(uncompressedData), dli)

			vii, err := rd.VarInt(&packetId)
			if err != nil {
				return nil, err
			}

			fmt.Println(packetId, vii)
			length -= int32(vii)
			data = make([]byte, length)
			if err := rd.FixedByteArray(data); err != nil {
				return nil, err
			}

			rd = io.NewReader(bytes.NewReader(data), int(length))
		}
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

func (conn *Conn) writeClassicDisconnect(reason string) {
	var data = []byte{0x0E}
	data = append(data, reason...)
	for len(data)-1 < 64 {
		data = append(data, 0x20)
	}

	conn.Write(data)
}

// returns true if the client is trying to log in to the server
func (conn *Conn) handleHandshake() bool {
	pk, err := conn.ReadPacket()
	if err != nil {
		conn.writeClassicDisconnect(clientVeryOldMsg)
		return false
	}
	handshaking, ok := pk.(*handshake.Handshaking)
	if !ok {
		if pk.ID() == 122 {
			conn.writeLegacyStatus(conn.listener.cfg.Status())
		}
		if pk.ID() == 78 {
			conn.writeLegacyDisconnect(clientTooOldMsg)
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
		if err := conn.WritePacket(&status.StatusResponse{Data: conn.listener.cfg.Status()}); err != nil {
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
		if handshaking.ProtocolVersion > ProtocolVersion {
			conn.WritePacket(&login.Disconnect{Reason: text.TextComponent{Text: clientTooNewMsg}})
			return false
		}
		if handshaking.ProtocolVersion < ProtocolVersion {
			conn.WritePacket(&login.Disconnect{Reason: text.TextComponent{Text: clientTooOldMsg}})
			return false
		}
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

		if conn.listener.cfg.Encrypt {
			if err := conn.encrypt(); err != nil {
				return false
			}
			if conn.listener.cfg.Authenticate {
				if err := conn.authenticate(); err != nil {
					conn.WritePacket(&login.Disconnect{Reason: text.TextComponent{Text: "This server uses authenticated encryption mode, and you are using a cracked account."}})
					return false
				}
			}
		}

		if err := conn.WritePacket(&login.SetCompression{Threshold: conn.listener.cfg.CompressionThreshold}); err != nil {
			return false
		}
		conn.compressionSet = true

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
