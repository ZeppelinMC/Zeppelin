// Package encoding provides encoding and decoding of minecraft data types
package encoding

import (
	"fmt"
	"io"
	"unsafe"
)

func AppendByte(data []byte, b int8) []byte {
	return append(data, byte(b))
}

func AppendUbyte(data []byte, b byte) []byte {
	return append(data, b)
}

func AppendShort(data []byte, s int16) []byte {
	return append(data, byte(s>>8), byte(s))
}

func AppendUshort(data []byte, s uint16) []byte {
	return append(data, byte(s>>8), byte(s))
}

func AppendInt(data []byte, i int32) []byte {
	return append(data, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

func AppendLong(data []byte, l int64) []byte {
	return append(data, byte(l>>56), byte(l>>48), byte(l>>40), byte(l>>32), byte(l>>24), byte(l>>16), byte(l>>8), byte(l))
}

func AppendVarInt(data []byte, value int32) []byte {
	ux := uint32(value)
	for ux >= 0x80 {
		data = append(data, byte(ux&0x7F)|0x80)
		ux >>= 7
	}
	return append(data, byte(ux))
}

func PutVarInt(data []byte, value int32) (n int) {
	ux := uint32(value)
	var i int
	for ; ux >= 0x80; i++ {
		data[i] = byte(ux&0x7F) | 0x80
		ux >>= 7
	}
	data[i] = byte(ux)

	return i
}

func WriteVarInt(w io.Writer, value int32) error {
	ux := uint32(value)
	for ux >= 0x80 {
		if _, err := w.Write([]byte{byte(ux&0x7F) | 0x80}); err != nil {
			return err
		}

		ux >>= 7
	}

	_, err := w.Write([]byte{byte(ux)})
	return err
}

func VarInt(data []byte) (int32, []byte, error) {
	var (
		position    int
		currentByte byte

		value int32
	)

	for {
		if len(data) == 0 {
			return value, data, io.EOF
		}
		currentByte = data[0]
		data = data[1:]

		value |= int32(currentByte&127) << position

		if (currentByte & 128) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return value, data, fmt.Errorf("VarInt is too big")
		}
	}

	return value, data, nil
}

func ReadVarInt(r io.Reader) (int32, error) {
	var (
		position    int32
		currentByte byte
		continueBit byte = 128
		segmentBits byte = 127

		value int32
	)

	for {
		if _, err := r.Read(unsafe.Slice(&currentByte, 1)); err != nil {
			return value, err
		}

		value |= int32(currentByte&segmentBits) << position

		if (currentByte & continueBit) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return value, fmt.Errorf("VarInt is too big")
		}
	}

	return value, nil
}

func AppendVarLong(data []byte, value int64) []byte {
	var (
		continueBit int64 = 128
		segmentBits int64 = 127
	)
	for {
		if (value & ^segmentBits) == 0 {
			return append(data, byte(value))
		}

		data = append(data, byte((value&segmentBits)|continueBit))

		value >>= 7
	}
}

func AppendString(data []byte, str string) []byte {
	data = AppendVarInt(data, int32(len(str)))

	return append(data, str...)
}

func String(data []byte) (string, error) {
	l, strd, err := VarInt(data)
	if err != nil {
		return "", err
	}

	return unsafe.String(unsafe.SliceData(strd), l), nil
}

type BitSet []int64

func (set BitSet) Get(i int) bool {
	return (set[i/64] & (1 << (i % 64))) != 0
}

func (set BitSet) Set(i int) {
	set[i/64] |= 1 << (i % 64)
}

func (set BitSet) Unset(i int) {
	set[i/64] &= ^(1 << (i % 64))
}

type FixedBitSet []byte

func (set FixedBitSet) Get(i int) bool {
	return (set[i/8] & (1 << (i % 8))) != 0
}

func (set FixedBitSet) Set(i int) {
	set[i/8] |= 1 << (i % 8)
}

func (set FixedBitSet) Unset(i int) {
	set[i/8] &= ^(1 << (i % 8))
}
