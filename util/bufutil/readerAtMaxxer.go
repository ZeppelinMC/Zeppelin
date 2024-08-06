package bufutil

import "io"

func NewReaderAtMaxxer(r io.ReaderAt, max int, baseOffset int64) *ReaderAtMaxxer {
	return &ReaderAtMaxxer{r: r, max: max, offset: baseOffset}
}

// ReaderAtMaxxer takes a readerAt and only allows max bytes to be read from it
type ReaderAtMaxxer struct {
	r io.ReaderAt

	max  int
	read int

	offset int64
}

func (r *ReaderAtMaxxer) Read(data []byte) (i int, err error) {
	if r.read >= r.max {
		return 0, io.EOF
	}

	remaining := r.max - r.read
	if len(data) > remaining {
		data = data[:remaining]
	}

	n, err := r.r.ReadAt(data, r.offset+int64(r.read))
	r.read += n

	if r.read >= r.max {
		err = io.EOF
	}

	return n, err
}
