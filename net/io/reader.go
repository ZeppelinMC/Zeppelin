package io

import (
	"aether/chat"
	"aether/nbt"
	"encoding/json"
	"fmt"
	"io"
	"unsafe"

	"github.com/google/uuid"
)

type Reader struct {
	r io.Reader
}

func NewReader(r io.Reader) Reader {
	return Reader{r}
}

func (r Reader) readBytes(l int) ([]byte, error) {
	arr := make([]byte, l)
	_, err := r.r.Read(arr)
	return arr, err
}

func (r Reader) Bool(b *bool) error {
	return r.Ubyte((*byte)(unsafe.Pointer(&b)))
}

func (r Reader) Byte(i *int8) error {
	d, err := r.readBytes(1)
	*i = int8(d[0])
	return err
}
func (r Reader) Ubyte(i *byte) error {
	d, err := r.readBytes(1)
	*i = d[0]
	return err
}

func (r Reader) Short(i *int16) error {
	d, err := r.readBytes(2)
	*i = int16(d[0])<<8 | int16(d[1])
	return err
}
func (r Reader) Ushort(i *uint16) error {
	d, err := r.readBytes(2)
	*i = uint16(d[0])<<8 | uint16(d[1])
	return err
}

func (r Reader) Int(i *int32) error {
	d, err := r.readBytes(4)
	*i = int32(d[0])<<24 | int32(d[1])<<16 | int32(d[2])<<8 | int32(d[3])
	return err
}

func (r Reader) Long(i *int64) error {
	d, err := r.readBytes(8)
	*i = int64(d[0])<<56 | int64(d[1])<<48 | int64(d[2])<<40 | int64(d[3])<<32 | int64(d[4])<<24 | int64(d[5])<<16 | int64(d[6])<<8 | int64(d[7])
	return err
}

func (r Reader) Float(f *float32) error {
	return r.Int((*int32)(unsafe.Pointer(f)))
}
func (r Reader) Double(f *float64) error {
	return r.Long((*int64)(unsafe.Pointer(f)))
}

func (r Reader) String(s *string) error {
	var l int32
	if _, err := r.VarInt(&l); err != nil {
		return err
	}
	d, err := r.readBytes(int(l))
	*s = *(*string)(unsafe.Pointer(&d))

	return err
}

func (r Reader) Identifier(s *string) error {
	return r.String(s)
}

func (r Reader) VarInt(value *int32) (i int, err error) {
	var (
		position     int32
		currentByte  byte
		CONTINUE_BIT byte = 128
		SEGMENT_BITS byte = 127

		size int
	)

	for {
		if err := r.Ubyte(&currentByte); err != nil {
			return size, err
		}
		*value |= int32((currentByte & SEGMENT_BITS)) << position
		size++

		if (currentByte & CONTINUE_BIT) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return size, fmt.Errorf("VarInt is too big")
		}
	}

	return size, nil
}

func (r Reader) VarLong(value *int64) error {
	var (
		position     int64
		currentByte  byte
		CONTINUE_BIT byte = 128
		SEGMENT_BITS byte = 127
	)

	for {
		if err := r.Ubyte(&currentByte); err != nil {
			return err
		}
		*value |= int64((currentByte & SEGMENT_BITS)) << position

		if (currentByte & CONTINUE_BIT) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return fmt.Errorf("VarInt is too big")
		}
	}

	return nil
}

func (r Reader) Position(x, y, z *int32) error {
	var l int64
	if err := r.Long(&l); err != nil {
		return err
	}

	*x = int32(l >> 38)
	*y = int32(l & ((1 << 12) - 1))
	*z = int32((l >> 12) & ((1 << 26) - 1))

	return nil
}

func (r Reader) UUID(u *uuid.UUID) error {
	d, err := r.readBytes(16)

	*u = uuid.UUID(d)
	return err
}

func (r Reader) BitSet(data *[]int64) error {
	var l int32
	if _, err := r.VarInt(&l); err != nil {
		return err
	}
	*data = make([]int64, l)

	for _, l := range *data {
		if err := r.Long(&l); err != nil {
			return err
		}
	}
	return nil
}

func (r Reader) ByteArray(s *[]byte) error {
	var l int32
	if _, err := r.VarInt(&l); err != nil {
		return err
	}
	d, err := r.readBytes(int(l))
	*s = d

	return err
}

func (r Reader) FixedByteArray(s []byte) error {
	_, err := r.r.Read(s)
	return err
}

func (r Reader) NBT(v any) error {
	dec := nbt.NewDecoder(r.r)
	dec.ReadRootName(false)

	_, err := dec.Decode(v)

	return err
}

func (r Reader) JSONTextComponent(comp *chat.TextComponent) error {
	var d []byte
	if err := r.ByteArray(&d); err != nil {
		return err
	}
	return json.Unmarshal(d, comp)
}
