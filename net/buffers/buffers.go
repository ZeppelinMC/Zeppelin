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
