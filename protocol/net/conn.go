package net

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"unicode/utf16"

	"github.com/zeppelinmc/zeppelin/protocol/net/cfb8"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/compress"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/handshake"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/login"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/status"
	"github.com/zeppelinmc/zeppelin/protocol/text"

	"github.com/google/uuid"
)

var PacketEncodeInterceptor func(c *Conn, pk packet.Encodeable) (stop bool)
var PacketDecodeInterceptor func(c *Conn, pk packet.Decodeable) (stop bool)

var PacketWriteInterceptor func(c *Conn, pk *bytes.Buffer, headerSize int32) (stop bool)
var PacketReadInterceptor func(c *Conn, pk *bytes.Reader, packetId int32) (stop bool)

const (
	clientVeryOldMsg = "Your client is WAYYYYYY too old!!! this server supports MC 1.21"
	clientTooOldMsg  = "Your client is too old! this server supports MC 1.21"
	clientTooNewMsg  = "Your client is too new! this server supports MC 1.21"
)

type Conn struct {
	net.Conn

	listener *Listener

	username   string
	uuid       uuid.UUID
	properties []login.Property

	state atomic.Int32

	encrypted                 bool
	sharedSecret, verifyToken []byte

	decrypter, encrypter *cfb8.CFB8
	compressionSet       bool

	usesForge bool
	read_mu   sync.Mutex
	write_mu  sync.Mutex
}

func (conn *Conn) UsesForge() bool {
	return conn.usesForge
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

var pkpool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(nil)
	},
}

var pkcppool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(nil)
	},
}

func (conn *Conn) WritePacket(pk packet.Encodeable) error {
	conn.write_mu.Lock() //TODO: fix concurrent ticking with encryption!
	defer conn.write_mu.Unlock()
	if PacketEncodeInterceptor != nil {
		if PacketEncodeInterceptor(conn, pk) {
			return nil
		}
	}

	var packetBuf = pkpool.Get().(*bytes.Buffer)
	packetBuf.Reset()
	defer pkpool.Put(packetBuf)

	w := encoding.NewWriter(packetBuf)
	// write the header for the packet

	var headerSize int32
	if conn.listener.cfg.CompressionThreshold < 0 || !conn.compressionSet {
		packetBuf.Write([]byte{0x80, 0x80, 0})
		headerSize = 3
	} else if conn.compressionSet {
		packetBuf.Write([]byte{0x80, 0x80, 0, 0x80, 0x80, 0})
		headerSize = 6
	}

	if err := w.VarInt(pk.ID()); err != nil {
		return err
	}
	if err := pk.Encode(w); err != nil {
		return err
	}

	if PacketWriteInterceptor != nil {
		if PacketWriteInterceptor(conn, packetBuf, headerSize) {
			return nil
		}
	}

	if conn.listener.cfg.CompressionThreshold < 0 || !conn.compressionSet { // no compression
		i := encoding.PutVarInt(packetBuf.Bytes()[:3], int32(packetBuf.Len()-3))
		if i != 2 {
			packetBuf.Bytes()[i] |= 0x80
		}

		_, err := packetBuf.WriteTo(conn)
		return err
	} else { // yes compression
		if conn.listener.cfg.CompressionThreshold > int32(packetBuf.Len())-6 { // packet is too small to be compressed
			i := encoding.PutVarInt(packetBuf.Bytes()[:3], int32(packetBuf.Len()-6))
			if i != 2 {
				packetBuf.Bytes()[i] |= 0x80
			}

			_, err := packetBuf.WriteTo(conn)
			return err
		} else { // packet is compressed
			uncompressedLength := int32(packetBuf.Len() - 6)
			if i := encoding.PutVarInt(packetBuf.Bytes()[3:6], uncompressedLength); i != 2 {
				packetBuf.Bytes()[i+3] |= 0x80
			}

			compressedPacket, err := compress.CompressZlib(packetBuf.Bytes()[6:], &MaxCompressedPacketSize)
			if err != nil {
				return err
			}

			packetBuf.Truncate(6)
			packetBuf.Write(compressedPacket)

			if i := encoding.PutVarInt(packetBuf.Bytes()[:3], int32(packetBuf.Len())-3); i != 2 {
				packetBuf.Bytes()[i] |= 0x80
			}

			if _, err := packetBuf.WriteTo(conn); err != nil {
				return err
			}

			return err
		}
	}
}

// 1MiB
var MaxCompressedPacketSize = 1024 * 1024

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

	return conn.Conn.Write(data)
}

func (conn *Conn) ReadPacket() (packet.Decodeable, error) {
	conn.read_mu.Lock()
	defer conn.read_mu.Unlock()

	var rd = encoding.NewReader(conn, 0)
	var (
		length, packetId int32
	)
	if conn.listener.cfg.CompressionThreshold < 0 || !conn.compressionSet { // no compression
		if _, err := rd.VarInt(&length); err != nil {
			return nil, err
		}
		if length <= 0 {
			return nil, fmt.Errorf("malformed packet: empty")
		}
		if length > 4096 {
			return nil, fmt.Errorf("packet too big")
		}

		var packet = make([]byte, length)
		if _, err := conn.Read(packet); err != nil {
			return nil, err
		}
		id, data, err := encoding.VarInt(packet)
		if err != nil {
			return nil, err
		}
		packetId = id
		packet = data
		length = int32(len(data))
		br := bytes.NewReader(packet)

		if PacketReadInterceptor != nil {
			if PacketReadInterceptor(conn, br, packetId) {
				return nil, fmt.Errorf("stopped by interceptor")
			}
		}

		rd = encoding.NewReader(br, int(length))
	} else {
		var packetLength int32
		if _, err := rd.VarInt(&packetLength); err != nil {
			return nil, err
		}
		if packetLength <= 0 {
			return nil, fmt.Errorf("malformed packet: empty")
		}
		var dataLength int32
		dataLengthSize, err := rd.VarInt(&dataLength)
		if err != nil {
			return nil, err
		}
		if dataLength < 0 {
			return nil, fmt.Errorf("malformed packet: negative length")
		}
		if dataLength == 0 { //packet is uncompressed
			length = packetLength - int32(dataLengthSize)
			if length < 0 {
				return nil, fmt.Errorf("malformed packet: negative length")
			}
			if length != 0 {
				var packet = make([]byte, length)
				if _, err := conn.Read(packet); err != nil {
					return nil, err
				}

				id, data, err := encoding.VarInt(packet)
				if err != nil {
					return nil, err
				}
				packetId = id
				packet = data
				length = int32(len(data))

				r := bytes.NewReader(packet)

				if PacketReadInterceptor != nil {
					if PacketReadInterceptor(conn, r, packetId) {
						return nil, fmt.Errorf("stopped by interceptor")
					}
				}

				rd = encoding.NewReader(r, int(length))
			}
		} else { //packet is compressed
			length = dataLength
			compressedLength := packetLength - int32(dataLengthSize)

			var ilength = int(length)

			var packetBuf = pkpool.Get().(*bytes.Buffer)
			packetBuf.Reset()
			packetBuf.ReadFrom(io.LimitReader(conn, int64(compressedLength)))
			defer pkpool.Put(packetBuf)

			uncompressedPacket, err := compress.DecompressZlib(packetBuf.Bytes(), &ilength)
			if err != nil {
				return nil, err
			}

			id, data, err := encoding.VarInt(uncompressedPacket)
			if err != nil {
				return nil, err
			}
			packetId = id
			uncompressedPacket = data
			length = int32(len(data))

			r := bytes.NewReader(uncompressedPacket)

			if PacketReadInterceptor != nil {
				if PacketReadInterceptor(conn, r, packetId) {
					return nil, fmt.Errorf("stopped by interceptor")
				}
			}

			rd = encoding.NewReader(r, int(length))
		}
	}

	var pk packet.Decodeable
	pc, ok := ServerboundPool[conn.state.Load()][packetId]

	if !ok {
		return packet.UnknownPacket{
			Id:      packetId,
			Length:  length,
			Payload: rd,
		}, nil
	} else {
		pk = pc()

		if PacketDecodeInterceptor != nil {
			if PacketDecodeInterceptor(conn, pk) {
				return nil, fmt.Errorf("stopped by interceptor")
			}
		}

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

var fml2hs = [...]byte{0, 70, 79, 82, 71, 69}

// Handles the handshake, and status/login state for the connection
// Returns true if the client is logging in
func (conn *Conn) handleHandshake() bool {
	pk, err := conn.ReadPacket()
	if err != nil {
		conn.writeClassicDisconnect(clientVeryOldMsg)
		return false
	}
	handshaking, ok := pk.(*handshake.Handshaking)
	if !ok {
		if pk.ID() == 122 {
			conn.writeLegacyStatus(conn.listener.cfg.Status(conn))
		}
		if pk.ID() == 78 {
			conn.writeLegacyDisconnect(clientTooOldMsg)
		}
		return false
	}
	if addr := []byte(handshaking.ServerAddress); len(addr) > len(fml2hs) && [6]byte(addr[len(addr)-len(fml2hs):]) == fml2hs {
		conn.usesForge = true
	}

	switch handshaking.NextState {
	case handshake.Status:
		conn.state.Store(StatusState)
		pk, err := conn.ReadPacket()
		if err != nil {
			return false
		}
		switch pac := pk.(type) {
		case *status.StatusRequest:
			if err := conn.WritePacket(&status.StatusResponse{Data: conn.listener.cfg.Status(conn)}); err != nil {
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

			conn.WritePacket(p)
		case *status.Ping:
			conn.WritePacket(pac)
		}
	case handshake.Transfer:
		if !conn.listener.cfg.AcceptTransfers {
			conn.WritePacket(&login.Disconnect{Reason: text.Sprint("Transfers are not allowed.")})
			conn.Close()
			return false
		}
		fallthrough
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
		if ok, r := conn.listener.ApprovePlayer(conn); !ok {
			var reason = text.Sprint("Disconnected")
			if r != nil {
				reason = *r
			}
			conn.WritePacket(&login.Disconnect{Reason: reason})
			conn.Close()
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
		if !conn.listener.cfg.Authenticate {
			conn.uuid = username2v3(conn.username)
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

func username2v3(username string) uuid.UUID {
	d := append([]byte("OfflinePlayer:"), username...)

	sum := md5.Sum(d)
	sum[6] &= 0x0f /* clear version        */
	sum[6] |= 0x30 /* set to version 3     */
	sum[8] &= 0x3f /* clear variant        */
	sum[8] |= 0x80 /* set to IETF variant  */

	return uuid.UUID(sum)
}
