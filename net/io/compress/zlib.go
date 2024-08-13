package compress

import (
	"bytes"
	"io"

	"github.com/4kills/go-libdeflate/v2"
	"github.com/zeppelinmc/zeppelin/net/io/buffers"
)

// Decompress zlib. If decompressedLength is provided, the data returned will only be safe to use until the next operation
// It is recommmended to provide the decompressed length to avoid allocation. If you don't know it, provide a number likely bigger than the decompressed length.
func DecompressZlib(data io.Reader, compressedLength int, decompressedLength *int) ([]byte, error) {
	var compressedBuffer = buffers.Buffers.Get().(*bytes.Buffer)
	defer buffers.Buffers.Put(compressedBuffer)
	compressedBuffer.Reset()
	if compressedBuffer.Len() < compressedLength {
		compressedBuffer.Grow(compressedLength)
	}

	if _, err := data.Read(compressedBuffer.Bytes()[:compressedLength]); err != nil {
		return nil, err
	}

	compressed := compressedBuffer.Bytes()[:compressedLength]

	dc := decompressors.Get().(libdeflate.Decompressor)
	defer decompressors.Put(dc)

	if decompressedLength != nil {
		dst := bufs.Get().([]byte)
		if len(dst) < int(*decompressedLength) {
			dst = make([]byte, *decompressedLength)
		}
		defer bufs.Put(dst)

		_, decompressedResult, err := dc.DecompressZlib(compressed, dst[:*decompressedLength])

		return decompressedResult, err
	} else {
		_, decompressedResult, err := dc.DecompressZlib(compressed, nil)
		return decompressedResult, err
	}
}

// Compresses zlib. If compressedLength is provided, data returned will only be safe to use until the next operation.
// It is recommmended to provide the compressed length to avoid allocation. If you don't know it, provide a number likely bigger than the compressed length.
func CompressZlib(decompressedData io.Reader, decompressedLength int, compressedLength *int) (compressed []byte, err error) {
	var decompressedBuffer = buffers.Buffers.Get().(*bytes.Buffer)
	defer buffers.Buffers.Put(decompressedBuffer)
	decompressedBuffer.Reset()
	if decompressedBuffer.Len() < decompressedLength {
		decompressedBuffer.Grow(decompressedLength)
	}

	if _, err := decompressedData.Read(decompressedBuffer.Bytes()[:decompressedLength]); err != nil {
		return nil, err
	}

	decompressed := decompressedBuffer.Bytes()[:decompressedLength]

	c := compressors.Get().(libdeflate.Compressor)
	defer compressors.Put(c)

	if compressedLength != nil {
		dst := bufs.Get().([]byte)
		if len(dst) < int(*compressedLength) {
			dst = make([]byte, *compressedLength)
		}
		defer bufs.Put(dst)

		_, compressedResult, err := c.CompressZlib(decompressed, dst[:*compressedLength])

		return compressedResult, err
	} else {
		_, compressedResult, err := c.CompressZlib(decompressed, nil)

		return compressedResult, err
	}
}
