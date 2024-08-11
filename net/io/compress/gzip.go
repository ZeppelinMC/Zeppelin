package compress

import (
	"bytes"
	"io"

	"github.com/4kills/go-libdeflate/v2"
	"github.com/zeppelinmc/zeppelin/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/net/io/util"
)

// Decompress gunzip. The data returned is only safe to use until the next operation
func DecompressGzip(data io.Reader, compressedLength int, decompressedLength *int) ([]byte, error) {
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
func CompressGzip(decompressedData io.Reader, decompressedLength int, compressedLength *int) (compressed []byte, err error) {
	var decompressedBuffer = buffers.Buffers.Get().(*bytes.Buffer)
	defer buffers.Buffers.Put(decompressedBuffer)
	decompressedBuffer.Reset()
	if decompressedBuffer.Len() < decompressedLength {
		decompressedBuffer.Grow(decompressedLength)
	}

	if _, err := decompressedBuffer.ReadFrom(util.NewReaderMaxxer(decompressedData, decompressedLength)); err != nil {
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

		_, compressedResult, err := c.Compress(decompressed, dst[:*compressedLength], libdeflate.ModeGzip)

		return compressedResult, err
	} else {
		_, compressedResult, err := c.Compress(decompressed, nil, libdeflate.ModeGzip)

		return compressedResult, err
	}
}
