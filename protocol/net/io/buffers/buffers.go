package buffers

import (
	"bytes"
	"sync"
)

var Buffers = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func Size() int {
	buf := Buffers.Get().(*bytes.Buffer)
	defer Buffers.Put(buf)

	return buf.Cap()
}
