package compress

import (
	"io"

	"github.com/4kills/go-libdeflate/v2"
)

// Decompress zlib. The data returned is only safe to use until the next operation
func DecompressZlib(data io.Reader, compressedLength int, decompressedLength *int) ([]byte, error) {
	var compressed = compressedBuffers.Get().([]byte)
	if len(compressed) < int(compressedLength) {
		compressed = make([]byte, compressedLength)
	}
	defer compressedBuffers.Put(compressed)

	if _, err := data.Read(compressed[:compressedLength]); err != nil {
		return nil, err
	}

	dc, err := libdeflate.NewDecompressor()
	if err != nil {
		return nil, err
	}

	if decompressedLength != nil {
		dst := decompressedBuffers.Get().([]byte)
		if len(dst) < int(*decompressedLength) {
			dst = make([]byte, *decompressedLength)
		}
		defer decompressedBuffers.Put(dst)

		_, decompressedResult, err := dc.DecompressZlib(compressed[:compressedLength], dst[:*decompressedLength])

		return decompressedResult, err
	} else {
		_, decompressedResult, err := dc.DecompressZlib(compressed[:compressedLength], nil)
		return decompressedResult, err
	}
}

func CompressZlib(decompressedData io.Reader, decompressedLength int, compressedLength *int) (compressed []byte, err error) {
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

		_, compressedResult, err := c.CompressZlib(decompressed[:decompressedLength], dst[:*compressedLength])

		return compressedResult, err
	} else {
		_, compressedResult, err := c.CompressZlib(decompressed[:decompressedLength], nil)

		return compressedResult, err
	}
}
