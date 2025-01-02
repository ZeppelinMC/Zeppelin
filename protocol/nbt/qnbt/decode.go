// deprecated: :(
package qnbt

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"unsafe"

	"github.com/oq-x/unsafe2"
)

type MismatchError struct {
	// name of the field trying to decode
	name string
	// the type kind expected
	expectedKind string
	// the type kind got
	gotKind string
	// the tag got
	tag string
}

func (m MismatchError) Error() string {
	return fmt.Sprintf("MismatchError: tag %s (expected %s), got %s (while decoding %s)", m.tag, m.expectedKind, m.gotKind, m.name)
}

var (
	ErrExpectPointer = errors.New("expected a pointer for dst")
	ErrUnsupported   = errors.New("unsupported type")
	ErrUnknownTag    = errors.New("unknown tag")
)

var decoders = sync.Pool{
	New: func() any {
		return new(Decoder)
	},
}

var readers = sync.Pool{
	New: func() any {
		r := new(bufferedReader)
		r.buf = make([]byte, 1024)
		return r
	},
}

type Decoder struct {
	// if disallowUnknownKeys is enabled, the decoder will return an error if it reads a name not in the struct
	disallowUnknownKeys bool
	// if networkNBT is enabled, the name for the root compound will not be read
	networkNBT bool

	rd *bufferedReader
}

// Close closes the decoder for further use and puts it back in the decoder pool
func (d *Decoder) Close() {
	readers.Put(d.rd)
	d.rd = nil
	decoders.Put(d)
}

func NewDecoder(rd io.Reader) *Decoder {
	dc := decoders.Get().(*Decoder)
	dc.disallowUnknownKeys, dc.networkNBT = false, false
	dc.rd = &bufferedReader{
		src: rd,
		buf: make([]byte, 1024),
	}

	dc.rd = readers.Get().(*bufferedReader)
	dc.rd.buf, dc.rd.src, dc.rd.r, dc.rd.w = slicesize(dc.rd.buf, 1024), rd, 0, 0

	return dc
}

func slicesize(s []byte, l int) []byte {
	if l < len(s) {
		return s[:l]
	} else {
		return make([]byte, l)
	}
}

func (dc *Decoder) DisallowUnknownKeys(v bool) *Decoder {
	dc.disallowUnknownKeys = v
	return dc
}

func (dc *Decoder) NetworkNBT(v bool) *Decoder {
	dc.networkNBT = v
	return dc
}

func Unmarshal(data []byte, dst any) (rootName string, err error) {
	dc := decoders.Get().(*Decoder)
	defer decoders.Put(dc)

	dc.disallowUnknownKeys, dc.networkNBT = false, false

	dc.rd = readers.Get().(*bufferedReader)
	dc.rd.buf, dc.rd.src, dc.rd.r, dc.rd.w = data, nil, 0, len(data)

	return dc.DecodeAndClose(dst)
}

// DecodeAndClose decodes into dst and closes the decoder
func (d *Decoder) DecodeAndClose(dst any) (rootName string, err error) {
	defer d.Close()
	return d.Decode(dst)
}

func (d *Decoder) Decode(dst any) (rootName string, err error) {
	id := unsafe2.InterfaceData(dst)
	typ := (*unsafe2.Type)(unsafe.Pointer(id.Type))

	if kindOf(typ) != pointer_k {
		return rootName, ErrExpectPointer
	}

	typ, v := ptrelem(typ, id.Value)

	rootTagS, err := d.rd.readBytesString(1)
	if err != nil {
		return rootName, err
	}

	if !d.networkNBT {
		if err := d.readString(&rootName); err != nil {
			return rootName, err
		}
	}

	switch rootTagS {
	case string_t_b:
		if k := kindOf(typ); k != string_k {
			return rootName, MismatchError{
				expectedKind: "string",
				gotKind:      kind_name(k),
				tag:          tagName(rootTagS),
				name:         "root",
			}
		}
	case compound_t_b:
		switch k := kindOf(typ); k {
		case struct_k:
			typ := (*structType)(unsafe.Pointer(typ))
			s := newStruct(typ, v)
			defer structs.Put(s)

			if err := d.decodeStructCompound(s); err != nil {
				return rootName, err
			}
		default:
			return rootName, MismatchError{
				expectedKind: "struct",
				gotKind:      kind_name(k),
				tag:          tagName(rootTagS),
				name:         "root",
			}
		}
	default:
		return rootName, ErrUnsupported
	}

	return rootName, nil
}

func (d *Decoder) discardCompound() error {
	for {
		tag, err := d.readByte2()
		if err != nil {
			return err
		}

		if tag == end_t {
			return nil
		}
		if err := d.skipString(); err != nil {
			return err
		}

		s, err := d.discard(tag)
		if err != nil {
			return err
		}
		if !s {
			switch tag {
			case compound_t:
				if err := d.discardCompound(); err != nil {
					return err
				}
			case list_t:
				if err := d.discardList(); err != nil {
					return err
				}
			default:
				return ErrUnknownTag
			}
		}
	}
}

func (d *Decoder) discard(tag_b byte) (skipped bool, err error) {
	switch tag_b {
	case byte_t:
		return true, d.rd.skip(1)
	case short_t:
		return true, d.rd.skip(2)
	case int_t, float_t:
		return true, d.rd.skip(4)
	case long_t, double_t:
		return true, d.rd.skip(8)
	case byte_array_t:
		l, err := d.readInt()
		if err != nil {
			return true, err
		}
		return true, d.rd.skip(int(l))
	case string_t:
		l, err := d.readShort()
		if err != nil {
			return true, err
		}
		return true, d.rd.skip(int(l))
	case int_array_t:
		l, err := d.readInt()
		if err != nil {
			return true, err
		}
		return true, d.rd.skip(int(l) * 4)
	case long_array_t:
		l, err := d.readInt()
		if err != nil {
			return true, err
		}
		return true, d.rd.skip(int(l) * 8)
	default:
		return
	}
}
