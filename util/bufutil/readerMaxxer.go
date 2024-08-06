package bufutil

import "io"

func NewReaderMaxxer(r io.Reader, max int) *ReaderMaxxer {
	return &ReaderMaxxer{r: r, max: max}
}

// ReaderMaxxer takes a reader and only allows max bytes to be read from it
type ReaderMaxxer struct {
	r io.Reader

	max  int
	read int
}

func (r *ReaderMaxxer) Read(data []byte) (i int, err error) {
	if r.read >= r.max {
		return 0, io.EOF
	}

	remaining := r.max - r.read
	if len(data) > remaining {
		data = data[:remaining]
	}

	n, err := r.r.Read(data)
	r.read += n

	if r.read >= r.max {
		err = io.EOF
	}

	return n, err
}
