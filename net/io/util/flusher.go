package util

import (
	"bytes"
	"io"
)

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{w: w}
}

type Flusher struct {
	w io.Writer

	buf bytes.Buffer
}

func (f *Flusher) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}

func (f *Flusher) Flush() (n int64, err error) {
	i, err := f.buf.WriteTo(f.w)

	return i, err
}
