package compress

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/pierrec/lz4/v4"
)

// Decompress a lz4-java data. The data returned is only safe to use until the next operation
func DecompressLZ4(data []byte) ([]byte, error) {
	var buf = bufs.Get().(*bytes.Buffer)
	defer bufs.Put(buf)

	buf.Reset()

	var (
		offsetFile   int
		offsetBuffer int
	)
	for {
		if offsetFile == len(data) {
			return buf.Bytes()[:offsetBuffer], nil
		}
		header := data[offsetFile : offsetFile+21]
		offsetFile += 21

		if string(header[:8]) != magic {
			return buf.Bytes()[:offsetBuffer], fmt.Errorf("invalid magic value")
		}

		token := header[8]
		compressedLength := int(binary.LittleEndian.Uint32(header[9:13]))

		compressionMethod := token & 0xf0

		switch compressionMethod {
		case methodUncompressed:
			if buf.Cap()-offsetBuffer < compressedLength {
				buf.Grow(offsetBuffer + compressedLength)
			}

			copy(buf.Bytes()[offsetBuffer:offsetBuffer+compressedLength], data[offsetFile:offsetFile+compressedLength])
			offsetBuffer += compressedLength
		case methodLZ4:
			decompressedLength := int(binary.LittleEndian.Uint32(header[13:17]))

			if buf.Cap()-offsetBuffer < decompressedLength {
				buf.Grow(offsetBuffer + decompressedLength)
			}

			if _, err := lz4.UncompressBlock(data[offsetFile:offsetFile+compressedLength], buf.Bytes()[offsetBuffer:offsetBuffer+decompressedLength]); err != nil {
				return buf.Bytes()[:offsetBuffer], err
			}
			offsetBuffer += decompressedLength
		}
		offsetFile += compressedLength
	}
}

const magic = "LZ4Block"
const (
	methodUncompressed = 1 << (iota + 4)
	methodLZ4
)
