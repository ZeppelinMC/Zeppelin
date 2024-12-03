package net

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/compress"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"io"
	"sync"
)

// PacketHandler handles all packet-related operations
type PacketHandler struct {
	conn        *Conn
	reader      *encoding.Reader
	readMutex   *sync.Mutex
	writeMutex  *sync.Mutex
	compression CompressionConfig
}

type CompressionConfig struct {
	threshold     int32
	enabled       bool
	maxPacketSize int
}

// NewPacketHandler creates a new packet handler for the connection
func NewPacketHandler(conn *Conn) *PacketHandler {
	reader := encoding.NewReader(conn, 0)
	return &PacketHandler{
		conn:       conn,
		reader:     &reader, // Take the address of the reader
		readMutex:  &conn.read_mu,
		writeMutex: &conn.write_mu,
		compression: CompressionConfig{
			threshold:     conn.listener.cfg.CompressionThreshold,
			enabled:       conn.compressionSet,
			maxPacketSize: MaxCompressedPacketSize,
		},
	}
}

func (ph *PacketHandler) readAndDecodePacket() (packet.Decodeable, bool, error) {
	packetData, packetId, err := ph.readPacketData()
	if err != nil {
		return nil, false, err
	}

	if stopped := ph.handleInterception(packetData, packetId); stopped {
		return nil, true, nil
	}

	return ph.decodePacket(packetData, packetId)
}

func (ph *PacketHandler) readPacketData() ([]byte, int32, error) {
	if !ph.compression.enabled {
		return ph.readUncompressedPacket()
	}
	return ph.readCompressedPacket()
}

func (ph *PacketHandler) readUncompressedPacket() ([]byte, int32, error) {
	var length int32
	if _, err := ph.reader.VarInt(&length); err != nil {
		return nil, 0, err
	}

	if err := validatePacketLength(length); err != nil {
		return nil, 0, err
	}

	packet := make([]byte, length)
	if _, err := ph.conn.Read(packet); err != nil {
		return nil, 0, err
	}

	id, data, err := encoding.VarInt(packet)
	if err != nil {
		return nil, 0, err
	}

	return data, id, nil
}

func (ph *PacketHandler) readCompressedPacket() ([]byte, int32, error) {
	packetLength, dataLength, err := ph.readCompressedHeaders()
	if err != nil {
		return nil, 0, err
	}

	if dataLength == 0 {
		return ph.readUncompressedDataFromCompressedPacket(packetLength)
	}
	return ph.readCompressedData(packetLength, dataLength)
}

func (ph *PacketHandler) readUncompressedDataFromCompressedPacket(packetLength int32) ([]byte, int32, error) {
	// Calculate actual data length (removing the size of the data length field)
	dataLengthSize := encoding.VarIntSize(0) // Size of a zero VarInt
	length := packetLength - dataLengthSize

	if length < 0 {
		return nil, 0, fmt.Errorf("malformed packet: negative length")
	}

	if length == 0 {
		return nil, 0, nil
	}

	// Read the raw packet data
	packet := make([]byte, length)
	if _, err := ph.conn.Read(packet); err != nil {
		return nil, 0, err
	}

	// Extract packet ID and data
	id, data, err := encoding.VarInt(packet)
	if err != nil {
		return nil, 0, err
	}

	return data, id, nil
}

func (ph *PacketHandler) readCompressedHeaders() (packetLength, dataLength int32, err error) {
	if _, err := ph.reader.VarInt(&packetLength); err != nil {
		return 0, 0, err
	}

	if err := validatePacketLength(packetLength); err != nil {
		return 0, 0, err
	}

	_, err = ph.reader.VarInt(&dataLength)
	if err != nil {
		return 0, 0, err
	}

	if dataLength < 0 {
		return 0, 0, fmt.Errorf("malformed packet: negative length")
	}

	return packetLength, dataLength, nil
}

func (ph *PacketHandler) readCompressedData(packetLength, dataLength int32) ([]byte, int32, error) {
	compressedLength := packetLength - int32(encoding.VarIntSize(dataLength))

	packetBuf := getBufferFromPool()
	defer returnBufferToPool(packetBuf)

	packetBuf.Reset()
	if _, err := packetBuf.ReadFrom(io.LimitReader(ph.conn, int64(compressedLength))); err != nil {
		return nil, 0, err
	}

	// Create a temporary int variable for the decompression
	ilength := int(dataLength)
	uncompressedPacket, err := compress.DecompressZlib(packetBuf.Bytes(), &ilength)
	if err != nil {
		return nil, 0, err
	}

	id, data, err := encoding.VarInt(uncompressedPacket)
	if err != nil {
		return nil, 0, err
	}

	return data, id, nil
}

// WritePacket handles writing a packet to the connection
func (ph *PacketHandler) WritePacket(pk packet.Encodeable) error {
	if ph.shouldLockWrite() {
		ph.writeMutex.Lock()
		defer ph.writeMutex.Unlock()
	}
	return ph.writePacketInternal(pk)
}

func (ph *PacketHandler) writePacketInternal(pk packet.Encodeable) error {
	if stopped := ph.handleEncodeInterception(pk); stopped {
		return nil
	}

	packetBuf := getBufferFromPool()
	defer returnBufferToPool(packetBuf)

	// Get header size based on compression
	headerSize := int32(3) // Default uncompressed header size
	if ph.compression.enabled {
		headerSize = 6 // Compressed header size
	}

	if err := ph.writePacketToBuffer(packetBuf, pk); err != nil {
		return err
	}

	if stopped := ph.handleWriteInterception(packetBuf, headerSize); stopped {
		return nil
	}

	return ph.writeBufferToConnection(packetBuf)
}

// Helper functions for buffer pool management
var packetBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(nil)
	},
}

func getBufferFromPool() *bytes.Buffer {
	return packetBufferPool.Get().(*bytes.Buffer)
}

func returnBufferToPool(buf *bytes.Buffer) {
	packetBufferPool.Put(buf)
}

// Helper function for packet length validation
func validatePacketLength(length int32) error {
	if length <= 0 {
		return fmt.Errorf("malformed packet: empty")
	}
	if length > 4096 {
		return fmt.Errorf("packet too big")
	}
	return nil
}

// Helper for UUID generation from username (offline mode)
func GenerateOfflineUUID(username string) uuid.UUID {
	data := append([]byte("OfflinePlayer:"), username...)
	sum := md5.Sum(data)

	// Set version to 3
	sum[6] &= 0x0f
	sum[6] |= 0x30

	// Set variant to IETF
	sum[8] &= 0x3f
	sum[8] |= 0x80

	return uuid.UUID(sum)
}

func (ph *PacketHandler) handleInterception(data []byte, packetId int32) bool {
	if PacketReadInterceptor == nil {
		return false
	}

	reader := bytes.NewReader(data)
	return PacketReadInterceptor(ph.conn, reader, packetId)
}

func (ph *PacketHandler) handleEncodeInterception(pk packet.Encodeable) bool {
	if PacketEncodeInterceptor == nil {
		return false
	}
	return PacketEncodeInterceptor(ph.conn, pk)
}

func (ph *PacketHandler) handleWriteInterception(buf *bytes.Buffer, headerSize int32) bool {
	if PacketWriteInterceptor == nil {
		return false
	}
	return PacketWriteInterceptor(ph.conn, buf, headerSize)
}

func (ph *PacketHandler) decodePacket(data []byte, packetId int32) (packet.Decodeable, bool, error) {
	pc, ok := ServerboundPool[ph.conn.state.Load()][packetId]
	if !ok {
		return packet.UnknownPacket{
			Id:      packetId,
			Length:  int32(len(data)),
			Payload: encoding.NewReader(bytes.NewReader(data), len(data)),
		}, false, nil
	}

	pk := pc()

	// Handle decode interception
	if PacketDecodeInterceptor != nil {
		if PacketDecodeInterceptor(ph.conn, pk) {
			return nil, true, nil
		}
	}

	rd := encoding.NewReader(bytes.NewReader(data), len(data))
	err := pk.Decode(rd)
	return pk, false, err
}

func (ph *PacketHandler) writePacketToBuffer(packetBuf *bytes.Buffer, pk packet.Encodeable) error {
	w := encoding.NewWriter(packetBuf)

	// Write header placeholder
	if ph.compression.enabled {
		packetBuf.Write([]byte{0x80, 0x80, 0, 0x80, 0x80, 0}) // 6 bytes for compressed
	} else {
		packetBuf.Write([]byte{0x80, 0x80, 0}) // 3 bytes for uncompressed
	}

	// Write packet ID and data
	if err := w.VarInt(pk.ID()); err != nil {
		return err
	}
	if err := pk.Encode(w); err != nil {
		return err
	}

	return nil
}

func (ph *PacketHandler) writeBufferToConnection(packetBuf *bytes.Buffer) error {
	if !ph.compression.enabled {
		return ph.writeUncompressedBuffer(packetBuf)
	}
	return ph.writeCompressedBuffer(packetBuf)
}

func (ph *PacketHandler) writeUncompressedBuffer(packetBuf *bytes.Buffer) error {
	// Write the packet length at the start
	i := encoding.PutVarInt(packetBuf.Bytes()[:3], int32(packetBuf.Len()-3))
	if i != 2 {
		packetBuf.Bytes()[i] |= 0x80
	}

	_, err := ph.conn.Write(packetBuf.Bytes())
	return err
}

func (ph *PacketHandler) writeCompressedBuffer(packetBuf *bytes.Buffer) error {
	dataLength := int32(packetBuf.Len() - 6) // Subtract header size

	// If packet is too small to compress
	if dataLength <= ph.compression.threshold {
		i := encoding.PutVarInt(packetBuf.Bytes()[:3], int32(packetBuf.Len()-3))
		if i != 2 {
			packetBuf.Bytes()[i] |= 0x80
		}
		_, err := ph.conn.Write(packetBuf.Bytes())
		return err
	}

	// Handle compression
	uncompressedLength := dataLength
	if i := encoding.PutVarInt(packetBuf.Bytes()[3:6], uncompressedLength); i != 2 {
		packetBuf.Bytes()[i+3] |= 0x80
	}

	compressedPacket, err := compress.CompressZlib(packetBuf.Bytes()[6:], &ph.compression.maxPacketSize)
	if err != nil {
		return err
	}

	// Reset buffer to header and write compressed data
	packetBuf.Truncate(6)
	packetBuf.Write(compressedPacket)

	// Write final packet length
	if i := encoding.PutVarInt(packetBuf.Bytes()[:3], int32(packetBuf.Len()-3)); i != 2 {
		packetBuf.Bytes()[i] |= 0x80
	}

	_, err = ph.conn.Write(packetBuf.Bytes())
	return err
}

func (ph *PacketHandler) shouldLockWrite() bool {
	// Only lock for encrypted connections
	return ph.conn.encrypted
}
