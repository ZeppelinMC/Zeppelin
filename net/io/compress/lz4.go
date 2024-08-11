package compress

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pierrec/lz4/v4"
	"github.com/zeppelinmc/zeppelin/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/net/io/util"
)

// Decompress an lz4-java block. The data returned is only safe to use until the next operation
func DecompressLZ4(data io.Reader) ([]byte, error) {
	var header [21]byte
	_, err := data.Read(header[:])
	if err != nil {
		return nil, err
	}
	magicValue := string(header[:8])
	if magicValue != magic {
		return nil, fmt.Errorf("invalid magic value")
	}

	compressedLength := int(binary.LittleEndian.Uint32(header[9:13]))
	decompressedLength := binary.LittleEndian.Uint32(header[13:17])

	var compressedBuffer = buffers.Buffers.Get().(*bytes.Buffer)
	defer buffers.Buffers.Put(compressedBuffer)
	compressedBuffer.Reset()
	if compressedBuffer.Len() < compressedLength {
		compressedBuffer.Grow(compressedLength)
	}

	if _, err := compressedBuffer.ReadFrom(util.NewReaderMaxxer(data, compressedLength)); err != nil {
		return nil, err
	}

	compressed := compressedBuffer.Bytes()[:compressedLength]

	token := header[8]
	compressionMethod := token & 0xf0
	switch compressionMethod {
	case methodLZ4:
		var decompressed = bufs.Get().([]byte)
		if len(decompressed) < int(decompressedLength) {
			decompressed = make([]byte, decompressedLength)
		}
		defer bufs.Put(decompressed)

		_, err = lz4.UncompressBlock(compressed, decompressed[:decompressedLength])

		return decompressed[:decompressedLength], err
	case methodUncompressed:
		return compressed, nil
	default:
		return nil, fmt.Errorf("unknown compression method %d", compressionMethod)
	}
}

const magic = "LZ4Block"
const (
	methodUncompressed = 1 << (iota + 4)
	methodLZ4
)
