package compress

import (
	"encoding/binary"
	"fmt"

	"github.com/pierrec/lz4/v4"
)

// Decompress an lz4-java block. The data returned is only safe to use until the next operation
func DecompressLZ4(data []byte) ([]byte, error) {
	if len(data) <= 21 {
		return nil, fmt.Errorf("missing header")
	}
	header := data[:21]
	data = data[21:]

	magicValue := string(header[:8])
	if magicValue != magic {
		return nil, fmt.Errorf("invalid magic value")
	}

	compressedLength := int(binary.LittleEndian.Uint32(header[9:13]))
	decompressedLength := binary.LittleEndian.Uint32(header[13:17])

	compressed := data[:compressedLength]

	token := header[8]
	compressionMethod := token & 0xf0
	switch compressionMethod {
	case methodLZ4:
		var decompressed = bufs.Get().([]byte)
		if len(decompressed) < int(decompressedLength) {
			decompressed = make([]byte, decompressedLength)
		}
		defer bufs.Put(decompressed)

		_, err := lz4.UncompressBlock(compressed, decompressed[:decompressedLength])

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
