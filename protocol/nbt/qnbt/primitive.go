package qnbt

import (
	"encoding/binary"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/protocol/nbt/qnbt/native"
)

const (
	end_t = iota
	byte_t
	short_t
	int_t
	long_t
	float_t
	double_t
	byte_array_t
	string_t
	list_t
	compound_t
	int_array_t
	long_array_t
)

// a little trick to avoid allocations
var (
	end_t_b        = string([]byte{end_t})
	byte_t_b       = string([]byte{byte_t})
	short_t_b      = string([]byte{short_t})
	int_t_b        = string([]byte{int_t})
	long_t_b       = string([]byte{long_t})
	float_t_b      = string([]byte{float_t})
	double_t_b     = string([]byte{double_t})
	byte_array_t_b = string([]byte{byte_array_t})
	string_t_b     = string([]byte{string_t})
	list_t_b       = string([]byte{list_t})
	compound_t_b   = string([]byte{compound_t})
	int_array_t_b  = string([]byte{int_array_t})
	long_array_t_b = string([]byte{long_array_t})
)

func tagName(t string) string {
	switch t {
	case end_t_b:
		return "TAG_End"
	case byte_t_b:
		return "TAG_Byte"
	case short_t_b:
		return "TAG_Short"
	case int_t_b:
		return "TAG_Int"
	case long_t_b:
		return "TAG_Long"
	case float_t_b:
		return "TAG_Float"
	case double_t_b:
		return "TAG_Double"
	case byte_array_t_b:
		return "TAG_Byte_Array"
	case string_t_b:
		return "TAG_String"
	case list_t_b:
		return "TAG_List"
	case compound_t_b:
		return "TAG_Compound"
	case int_array_t_b:
		return "TAG_Int_Array"
	case long_array_t_b:
		return "TAG_Long_Array"
	default:
		return "Unknown"
	}
}

func (d *Decoder) load(length int) []byte {
	b := d.rd.buf[d.rd.r : d.rd.r+length]
	d.rd.r += length

	return b
}

func (d *Decoder) readByte(b *byte) error {
	if err := d.rd.fill(1); err != nil {
		return err
	}

	*b = d.rd.buf[d.rd.r]
	d.rd.r++

	return nil
}

func (d *Decoder) readByte2() (byte, error) {
	if err := d.rd.fill(1); err != nil {
		return 0, err
	}

	b := d.rd.buf[d.rd.r]
	d.rd.r++

	return b, nil
}
func (d *Decoder) readBytePtr(ptr unsafe.Pointer) error {
	return d.readByte((*byte)(ptr))
}

func (d *Decoder) readShort() (int16, error) {
	if err := d.rd.fill(2); err != nil {
		return 0, err
	}

	return int16(binary.BigEndian.Uint16(d.load(2))), nil
}

func (d *Decoder) readShortPtr(ptr unsafe.Pointer) error {
	if err := d.rd.fill(2); err != nil {
		return err
	}
	s := unsafe.Slice((*byte)(ptr), 2)

	copy(s, d.load(2))

	if !native.SystemBigEndian {
		native.Convert16(s)
	}

	return nil
}

func (d *Decoder) readInt() (int32, error) {
	if err := d.rd.fill(4); err != nil {
		return 0, err
	}

	return int32(binary.BigEndian.Uint32(d.load(4))), nil
}

func (d *Decoder) readIntPtr(ptr unsafe.Pointer) error {
	if err := d.rd.fill(4); err != nil {
		return err
	}
	s := unsafe.Slice((*byte)(ptr), 4)

	copy(s, d.load(4))

	if !native.SystemBigEndian {
		native.Convert32(s)
	}

	return nil
}

func (d *Decoder) readLongPtr(ptr unsafe.Pointer) error {
	if err := d.rd.fill(8); err != nil {
		return err
	}
	s := unsafe.Slice((*byte)(ptr), 8)

	copy(s, d.load(8))

	if !native.SystemBigEndian {
		native.Convert64(s)
	}

	return nil
}

func (d *Decoder) skipString() error {
	l, err := d.readShort()
	if err != nil {
		return err
	}
	return d.rd.skip(int(l))
}

func (d *Decoder) readString(dst *string) error {
	l, err := d.readShort()
	if err != nil {
		return err
	}

	if err := d.rd.fill(int(l)); err != nil {
		return err
	}

	*dst = string(d.load(int(l)))

	return nil
}

func (d *Decoder) readStringNonCopy(dst *string) error {
	l, err := d.readShort()
	if err != nil {
		return err
	}

	if err := d.rd.fill(int(l)); err != nil {
		return err
	}

	buf := d.rd.buf[d.rd.r : d.rd.r+int(l)]
	d.rd.r += int(l)

	*dst = unsafe.String(unsafe.SliceData(buf), len(buf))

	return nil
}

func (d *Decoder) readStringPtr(dst unsafe.Pointer) error {
	return d.readString((*string)(dst))
}

func (d *Decoder) readLongArray(ptr unsafe.Pointer) error {
	l, err := d.readInt()
	if err != nil {
		return err
	}

	return d.decodeLongArray(int(l), ptr)
}

func (d *Decoder) decodeLongArray(len int, ptr unsafe.Pointer) error {
	for i := 0; i < len; i++ {
		if err := d.readLongPtr(ptr); err != nil {
			return err
		}

		ptr = unsafe.Add(ptr, 8)
	}

	return nil
}

func (d *Decoder) readIntArray(ptr unsafe.Pointer) error {
	l, err := d.readInt()
	if err != nil {
		return err
	}

	return d.decodeIntArray(int(l), ptr)
}

func (d *Decoder) decodeIntArray(len int, ptr unsafe.Pointer) error {
	for i := 0; i < len; i++ {
		if err := d.readIntPtr(ptr); err != nil {
			return err
		}

		ptr = unsafe.Add(ptr, 4)
	}

	return nil
}

func (d *Decoder) readByteArray(ptr unsafe.Pointer) error {
	l, err := d.readInt()
	if err != nil {
		return err
	}

	return d.decodeByteArray(int(l), ptr)
}

func (d *Decoder) decodeByteArray(len int, ptr unsafe.Pointer) error {
	data, err := d.rd.readBytes(len)
	if err != nil {
		return err
	}
	copy(unsafe.Slice((*byte)(ptr), len), data)

	return nil
}
