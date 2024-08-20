package compress

import (
	"bytes"
	"sync"

	"github.com/4kills/go-libdeflate/v2"
)

var decompressors = sync.Pool{
	New: func() any {
		dc, _ := libdeflate.NewDecompressor()
		return dc
	},
}

var compressors = sync.Pool{
	New: func() any {
		dc, _ := libdeflate.NewCompressor()
		return dc
	},
}

var bufs = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}
