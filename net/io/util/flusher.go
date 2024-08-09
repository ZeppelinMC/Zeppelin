package util

import (
	"bytes"
	"io"
)

func NewFlusher(w io.WriteCloser, buf *bytes.Buffer) *Flusher {
	f := &Flusher{w: w}
	if buf != nil {
		f.buf = *buf
	}

	return f
}

type Flusher struct {
	w io.WriteCloser

	buf bytes.Buffer
}

func (f *Flusher) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}

// flush writes the buffer to w and closes it
func (f *Flusher) Flush() (n int64, err error) {
	i, err := f.buf.WriteTo(f.w)

	if err != nil {
		return i, err
	}
	return i, f.w.Close()
}
