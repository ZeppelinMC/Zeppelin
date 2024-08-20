package compress

import (
	"github.com/4kills/go-libdeflate/v2"
)

// Decompress gunzip. The data returned is only safe to use until the next operation
func DecompressGzip(compressed []byte, decompressedLength *int) ([]byte, error) {
	dc := decompressors.Get().(libdeflate.Decompressor)
	defer decompressors.Put(dc)

	if decompressedLength != nil {
		dst := bufs.Get().([]byte)
		if len(dst) < int(*decompressedLength) {
			dst = make([]byte, *decompressedLength)
		}
		defer bufs.Put(dst)

		_, decompressedResult, err := dc.Decompress(compressed, dst[:*decompressedLength], libdeflate.ModeGzip)

		return decompressedResult, err
	} else {
		_, decompressedResult, err := dc.Decompress(compressed, nil, libdeflate.ModeGzip)
		return decompressedResult, err
	}
}

// Compresses gunzip. If compressedLength is provided, data returned will only be safe to use until the next operation.
// It is recommmended to provide the compressed length to avoid allocation. If you don't know it, provide a number likely bigger than the compressed length.
func CompressGzip(decompressedData []byte, compressedLength *int) (compressed []byte, err error) {
	c := compressors.Get().(libdeflate.Compressor)
	defer compressors.Put(c)

	if compressedLength != nil {
		dst := bufs.Get().([]byte)
		if len(dst) < int(*compressedLength) {
			dst = make([]byte, *compressedLength)
		}
		defer bufs.Put(dst)

		_, compressedResult, err := c.Compress(decompressedData, dst[:*compressedLength], libdeflate.ModeGzip)

		return compressedResult, err
	} else {
		_, compressedResult, err := c.Compress(decompressedData, nil, libdeflate.ModeGzip)

		return compressedResult, err
	}
}
