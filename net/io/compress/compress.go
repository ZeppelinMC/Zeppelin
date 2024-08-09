package compress

import "sync"

var decompressedBuffers = sync.Pool{
	New: func() any { return make([]byte, 0) },
}

var compressedBuffers = sync.Pool{
	New: func() any { return make([]byte, 0) },
}
