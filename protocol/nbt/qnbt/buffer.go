package qnbt

import (
	"io"
	"unsafe"
)

type bufferedReader struct {
	src io.Reader
	buf []byte

	r, w int
}

func (rd *bufferedReader) available() int {
	return rd.w - rd.r
}

func (rd *bufferedReader) readBytes(num int) ([]byte, error) {
	if err := rd.fill(num); err != nil {
		return nil, err
	}
	b := rd.buf[rd.r : rd.r+num]
	rd.r += num

	return b, nil
}

func (rd *bufferedReader) readBytesString(num int) (string, error) {
	b, err := rd.readBytes(num)
	if err != nil {
		return "", err
	}

	return unsafe.String(unsafe.SliceData(b), num), nil
}

func (rd *bufferedReader) fill(min int) error {
	if rd.available() < min {
		rd.r, rd.w = 0, copy(rd.buf, rd.buf[rd.r:rd.w])

		for i := 0; i < 5; i++ {
			n, err := rd.src.Read(rd.buf[rd.w:])
			rd.w += n
			if err != nil {
				return err
			}

			min -= n
			if min <= 0 || rd.w == len(rd.buf) {
				return nil
			}
		}
		return io.ErrNoProgress
	}

	return nil
}

func (rd *bufferedReader) skip(n int) error {
	avail := rd.w - rd.r
	if avail > n {
		rd.r += n
		return nil
	}

	n -= avail
	rd.r += avail

	rd.r, rd.w = 0, copy(rd.buf, rd.buf[rd.r:rd.w])

	for n > 0 {
		mx := n
		if mx >= len(rd.buf) {
			mx = len(rd.buf)
		}
		x, err := rd.src.Read(rd.buf[rd.w : rd.w+mx])
		rd.w += x
		rd.r += x
		if err != nil {
			return err
		}

		if rd.w == len(rd.buf) {
			rd.r, rd.w = 0, 0
		}

		n -= x
	}

	return nil
}
