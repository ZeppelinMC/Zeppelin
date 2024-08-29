package buffers

import (
	"bytes"
	"sync"
)

var Buffers = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(nil)
	},
}

func Size() int {
	buf := Buffers.Get().(*bytes.Buffer)
	defer Buffers.Put(buf)

	return buf.Len()
}

func Reset() {
	buf := Buffers.Get().(*bytes.Buffer)
	buf.Reset()
	Buffers.Put(buf)
}
