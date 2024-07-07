package io

import (
	"aether/chat"
	"aether/nbt"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"unsafe"

	"github.com/google/uuid"
)

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) Writer {
	return Writer{w}
}

func (w Writer) writeBytes(bytes ...byte) error {
	_, err := w.w.Write(bytes)
	return err
}

func (w Writer) Bool(b bool) error {
	return w.Ubyte(*(*byte)(unsafe.Pointer(&b)))
}

func (w Writer) Byte(i int8) error {
	return w.writeBytes(byte(i))
}
func (w Writer) Ubyte(i uint8) error {
	return w.writeBytes(i)
}

func (w Writer) Short(i int16) error {
	return w.writeBytes(
		byte(i>>8),
		byte(i),
	)
}
func (w Writer) Ushort(i uint16) error {
	return w.writeBytes(
		byte(i>>8),
		byte(i),
	)
}

func (w Writer) Int(i int32) error {
	return w.writeBytes(
		byte(i>>24),
		byte(i>>16),
		byte(i>>8),
		byte(i),
	)
}

func (w Writer) Long(i int64) error {
	return w.writeBytes(
		byte(i>>56),
		byte(i>>48),
		byte(i>>40),
		byte(i>>32),
		byte(i>>24),
		byte(i>>16),
		byte(i>>8),
		byte(i),
	)
}

func (w Writer) Float(f float32) error {
	return w.Int(int32(math.Float32bits(f)))
}
func (w Writer) Double(f float64) error {
	return w.Long(int64(math.Float64bits(f)))
}

func (w Writer) String(s string) error {
	if err := w.VarInt(int32(len(s))); err != nil {
		return err
	}
	return w.writeBytes(*(*[]byte)(unsafe.Pointer(&s))...)
}

func (w Writer) Identifier(s string) error {
	if len(s) > 32767 {
		return fmt.Errorf("expected identifier len to be > 32767, got %d", len(s))
	}
	if err := w.Int(int32(len(s))); err != nil {
		return err
	}
	return w.writeBytes(*(*[]byte)(unsafe.Pointer(&s))...)
}

func (w Writer) VarInt(value int32) error {
	var (
		CONTINUE_BIT int32 = 128
		SEGMENT_BITS int32 = 127
	)
	for {
		if (value & ^SEGMENT_BITS) == 0 {
			return w.Ubyte(byte(value))
		}

		if err := w.Ubyte(byte((value & SEGMENT_BITS) | CONTINUE_BIT)); err != nil {
			return err
		}

		value >>= 7
	}
}

func (w Writer) VarLong(value int64) error {
	var (
		CONTINUE_BIT int64 = 128
		SEGMENT_BITS int64 = 127
	)
	for {
		if (value & ^SEGMENT_BITS) == 0 {
			return w.Ubyte(byte(value))
		}

		if err := w.Ubyte(byte((value & SEGMENT_BITS) | CONTINUE_BIT)); err != nil {
			return err
		}

		value >>= 7
	}
}

func (w Writer) Position(x, y, z int32) error {
	return w.Long(((int64(x) & 0x3FFFFFF) << 38) | ((int64(z) & 0x3FFFFFF) << 12) | (int64(y) & 0xFFF))
}

func (w Writer) UUID(u uuid.UUID) error {
	d := [16]byte(u)
	return w.writeBytes(d[:]...)
}

func (w Writer) BitSet(data []int64) error {
	if err := w.VarInt(int32(len(data))); err != nil {
		return err
	}
	for _, l := range data {
		if err := w.Long(l); err != nil {
			return err
		}
	}
	return nil
}

func (w Writer) ByteArray(s []byte) error {
	if err := w.VarInt(int32(len(s))); err != nil {
		return err
	}
	return w.writeBytes(s...)
}

func (w Writer) FixedByteArray(s []byte) error {
	return w.writeBytes(s...)
}

func (w Writer) JSONTextComponent(comp chat.TextComponent) error {
	d, _ := json.Marshal(comp)

	return w.ByteArray(d)
}

func (w Writer) NBT(data any) error {
	enc := nbt.NewEncoder(w.w)
	enc.WriteRootName(false)

	return enc.Encode("", data)
}
