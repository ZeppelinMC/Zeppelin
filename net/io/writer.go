package io

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/nbt"
	"github.com/zeppelinmc/zeppelin/text"

	"github.com/google/uuid"
)

type Writer struct {
	w *bytes.Buffer
}

func NewWriter(w *bytes.Buffer) Writer {
	return Writer{w}
}

func (w Writer) Write(data []byte) (i int, err error) {
	return w.w.Write(data)
}

func (w Writer) Bool(b bool) error {
	return w.Ubyte(*(*byte)(unsafe.Pointer(&b)))
}

func (w Writer) Byte(i int8) error {
	return w.Ubyte(uint8(i))
}
func (w Writer) Ubyte(i uint8) error {
	_, err := w.Write([]byte{i})

	return err
}

func (w Writer) Short(i int16) error {
	return w.Ushort(uint16(i))
}
func (w Writer) Ushort(i uint16) error {
	_, err := w.Write(
		[]byte{
			byte(i >> 8),
			byte(i),
		},
	)
	return err
}

func (w Writer) Int(i int32) error {
	_, err := w.Write(
		[]byte{
			byte(i >> 24),
			byte(i >> 16),
			byte(i >> 8),
			byte(i),
		},
	)
	return err
}

func (w Writer) Long(i int64) error {
	_, err := w.Write(
		[]byte{
			byte(i >> 56),
			byte(i >> 48),
			byte(i >> 40),
			byte(i >> 32),
			byte(i >> 24),
			byte(i >> 16),
			byte(i >> 8),
			byte(i),
		},
	)
	return err
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
	_, err := w.Write(unsafe.Slice(unsafe.StringData(s), len(s))) //(*(*[]byte)(unsafe.Pointer(&s)))
	return err
}

func (w Writer) Identifier(s string) error {
	if len(s) > 32767 {
		return fmt.Errorf("expected identifier len to be smaller than 32767, got %d", len(s))
	}
	if err := w.VarInt(int32(len(s))); err != nil {
		return err
	}
	_, err := w.Write(unsafe.Slice(unsafe.StringData(s), len(s))) //(*(*[]byte)(unsafe.Pointer(&s)))
	return err
}

func (w Writer) VarInt(value int32) error {
	ux := uint32(value)
	for ux >= 0x80 {
		if err := w.Ubyte(byte(ux&0x7F) | 0x80); err != nil {
			return err
		}

		ux >>= 7
	}

	return w.Ubyte(byte(ux))
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
	_, err := w.Write(d[:])
	return err
}

func (w Writer) BitSet(data BitSet) error {
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

func (w Writer) FixedBitSet(data FixedBitSet) error {
	for _, l := range data {
		if err := w.Ubyte(l); err != nil {
			return err
		}
	}
	return nil
}

// Length prefixed byte array
func (w Writer) ByteArray(s []byte) error {
	if err := w.VarInt(int32(len(s))); err != nil {
		return err
	}
	_, err := w.w.Write(s)

	return err
}

func (w Writer) FixedByteArray(s []byte) error {
	_, err := w.w.Write(s)

	return err
}

func (w Writer) JSONTextComponent(comp text.TextComponent) error {
	d, _ := json.Marshal(comp)

	return w.ByteArray(d)
}

func (w Writer) TextComponent(comp text.TextComponent) error {
	return w.NBT(comp)
}

func (w Writer) StringTextComponent(text string) error {
	if err := w.Byte(8); err != nil {
		return err
	}
	if err := w.Short(int16(len(text))); err != nil {
		return err
	}

	return w.String(text)
}

func (w Writer) NBT(data any) error {
	enc := nbt.NewEncoder(w.w)
	enc.WriteRootName(false)

	return enc.Encode("", data)
}
