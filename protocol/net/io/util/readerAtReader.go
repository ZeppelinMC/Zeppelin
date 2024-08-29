package util

import "io"

type ReaderAtReader struct {
	r io.ReaderAt

	offset int64
}

// NewReaderAtReader makes a ReaderAtReader that starts reading at off
func NewReaderAtReader(r io.ReaderAt, off int64) *ReaderAtReader {
	return &ReaderAtReader{r: r}
}

func (r *ReaderAtReader) Read(p []byte) (n int, err error) {
	n, err = r.r.ReadAt(p, r.offset)
	if err == nil {
		r.offset += int64(n)
	}

	return n, err
}
