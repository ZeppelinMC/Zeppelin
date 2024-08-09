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
