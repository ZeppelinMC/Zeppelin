package compress

import (
	"bytes"

	"github.com/4kills/go-libdeflate/v2"
)

// Decompress zlib. If decompressedLength is provided, the data returned will only be safe to use until the next operation
// It is recommmended to provide the decompressed length to avoid allocation. If you don't know it, provide a number likely bigger than the decompressed length.
func DecompressZlib(compressed []byte, decompressedLength *int) ([]byte, error) {
	dc := decompressors.Get().(libdeflate.Decompressor)
	defer decompressors.Put(dc)

	if decompressedLength != nil {
		dst := bufs.Get().(*bytes.Buffer)
		if dst.Len() < int(*decompressedLength) {
			dst.Grow(*decompressedLength)
		}
		defer bufs.Put(dst)

		c := dst.Bytes()[:*decompressedLength]

		_, decompressedResult, err := dc.DecompressZlib(compressed, c)

		return decompressedResult, err
	} else {
		_, decompressedResult, err := dc.DecompressZlib(compressed, nil)
		return decompressedResult, err
	}
}

// Compresses zlib. If compressedLength is provided, data returned will only be safe to use until the next operation.
// It is recommmended to provide the compressed length to avoid allocation. If you don't know it, provide a number likely bigger than the compressed length.
func CompressZlib(decompressedData []byte, compressedLength *int) (compressed []byte, err error) {
	c := compressors.Get().(libdeflate.Compressor)
	defer compressors.Put(c)

	if compressedLength != nil {
		dst := bufs.Get().(*bytes.Buffer)
		if dst.Len() < int(*compressedLength) {
			dst.Grow(*compressedLength)
		}
		defer bufs.Put(dst)

		dc := dst.Bytes()[:*compressedLength]

		_, compressedResult, err := c.CompressZlib(decompressedData, dc)

		return compressedResult, err
	} else {
		_, compressedResult, err := c.CompressZlib(decompressedData, nil)

		return compressedResult, err
	}
}
