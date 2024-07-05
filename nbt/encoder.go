package nbt

import (
	"io"
	"math"
	"unsafe"
)

type Encoder struct {
	w io.Writer
}

func (e Encoder) writeBytes(b ...byte) error {
	_, err := e.w.Write(b)
	return err
}

func (e Encoder) writeByte(b int8) error {
	return e.writeBytes(byte(b))
}

func (e Encoder) writeShort(s int16) error {
	return e.writeBytes(
		byte(s>>8),
		byte(s),
	)
}

func (e Encoder) writeInt(i int32) error {
	return e.writeBytes(
		byte(i>>24),
		byte(i>>16),
		byte(i>>8),
		byte(i),
	)
}

func (e Encoder) writeLong(l int64) error {
	return e.writeBytes(
		byte(l>>56),
		byte(l>>48),
		byte(l>>40),
		byte(l>>32),
		byte(l>>24),
		byte(l>>16),
		byte(l>>8),
		byte(l),
	)
}

func (e Encoder) writeFloat(f float32) error {
	return e.writeInt(int32(math.Float32bits(f)))
}

func (e Encoder) writeDouble(d float64) error {
	return e.writeLong(int64(math.Float64bits(d)))
}

func (e Encoder) writeByteArray(ba []int8) error {
	if err := e.writeInt(int32(len(ba))); err != nil {
		return err
	}
	return e.writeBytes(*(*[]byte)(unsafe.Pointer(&ba))...)
}

func (e Encoder) writeIntList(il []int32) error {
	if err := e.writeInt(int32(len(il))); err != nil {
		return err
	}
	for _, i := range il {
		if err := e.writeInt(i); err != nil {
			return err
		}
	}
	return nil
}

func (e Encoder) writeLongList(il []int64) error {
	if err := e.writeInt(int32(len(il))); err != nil {
		return err
	}
	for _, i := range il {
		if err := e.writeLong(i); err != nil {
			return err
		}
	}
	return nil
}
