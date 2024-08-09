package compress

import (
	"io"

	"github.com/4kills/go-libdeflate/v2"
)

// Decompress gunzip. The data returned is only safe to use until the next operation
func DecompressGzip(data io.Reader, compressedLength int, decompressedLength *int) ([]byte, error) {
	var compressed = compressedBuffers.Get().([]byte)
	if len(compressed) < int(compressedLength) {
		compressed = make([]byte, compressedLength)
	}
	defer compressedBuffers.Put(compressed)

	var decompressed []byte
	if decompressedLength != nil {
		var decompressed = decompressedBuffers.Get().([]byte)
		if len(decompressed) < int(*decompressedLength) {
			decompressed = make([]byte, *decompressedLength)
		}
		defer decompressedBuffers.Put(decompressed)
	}

	dc, err := libdeflate.NewDecompressor()
	if err != nil {
		return nil, err
	}
	_, decompressedResult, err := dc.Decompress(compressed, decompressed, libdeflate.ModeGzip)

	return decompressedResult, err
}

// Compresses gunzip. If compressedLength is provided, data returned will only be safe to use until the next operation.
// It is recommmended to provide the compressed length to avoid allocation. If you don't know it, provide a number likely bigger than the compressed length.
func CompressGzip(decompressedData io.Reader, decompressedLength int, compressedLength *int) (compressed []byte, err error) {
	c, err := libdeflate.NewCompressor()
	if err != nil {
		return nil, err
	}

	var decompressed = decompressedBuffers.Get().([]byte)
	if len(decompressed) < int(decompressedLength) {
		decompressed = make([]byte, decompressedLength)
	}
	defer decompressedBuffers.Put(decompressed)

	if _, err := decompressedData.Read(decompressed[:decompressedLength]); err != nil {
		return nil, err
	}

	if compressedLength != nil {
		dst := compressedBuffers.Get().([]byte)
		if len(dst) < int(*compressedLength) {
			dst = make([]byte, *compressedLength)
		}
		defer compressedBuffers.Put(dst)

		_, compressedResult, err := c.Compress(decompressed[:decompressedLength], dst[:*compressedLength], libdeflate.ModeGzip)

		return compressedResult, err
	} else {
		_, compressedResult, err := c.Compress(decompressed[:decompressedLength], nil, libdeflate.ModeGzip)

		return compressedResult, err
	}
}
