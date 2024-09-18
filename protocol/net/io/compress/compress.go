package compress

import (
	"bytes"
	"compress/gzip"
	"sync"

	"github.com/4kills/go-libdeflate/v2"
	"github.com/4kills/go-zlib"
)

var Decompressors = sync.Pool{
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

// RZlib is a pool of zlib.Reader's that can be reused to avoid allocations
var RZlib = sync.Pool{
	New: func() any { d, _ := zlib.NewReader(nil); return d },
}

// RGzip is a pool of gzip.Reader's that can be reused to avoid allocations
var RGzip = sync.Pool{
	New: func() any { d, _ := gzip.NewReader(nil); return d },
}

// WZlib is a pool of zlib.Writer's that can be reused to avoid allocations
var WZlib = sync.Pool{
	New: func() any { return zlib.NewWriter(nil) },
}

// WGzip is a pool of gzip.Writer's that can be reused to avoid allocations
var WGzip = sync.Pool{
	New: func() any { return gzip.NewWriter(nil) },
}
