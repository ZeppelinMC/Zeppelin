package qnbt

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"

	"github.com/oq-x/unsafe2"
)

var structs = sync.Pool{
	New: func() any {
		return &struct_t{
			names: make(map[string]int),
		}
	}}

func newStruct(t *structType, ptr unsafe.Pointer) *struct_t {
	s := structs.Get().(*struct_t)
	clear(s.names)

	s.t = t
	s.ptr = ptr

	for i, f := range t.fields {
		if !f.name.exported() {
			continue
		}
		name := f.name.name()

		if nbtname, ok := findNBTTag(f.name.tag()); ok {
			name = nbtname
			if name == "-" {
				continue
			}
			if i := strings.Index(name, ",omitempty"); i != -1 {
				name = name[:i]
			}
		}

		s.names[name] = i
	}

	return s
}

func findNBTTag(tag string) (string, bool) {
	i := strings.Index(tag, `nbt:"`)
	if i == -1 {
		return "", false
	}
	tag = tag[i+5:]
	i = strings.Index(tag, `"`)
	if i == -1 {
		return "", false
	}

	return tag[:i], true
}

type struct_t struct {
	t *structType

	names map[string]int
	ptr   unsafe.Pointer
}

func (s *struct_t) field(n string) (f *structField, ptr unsafe.Pointer, ok bool) {
	i, ok := s.names[n]

	if !ok {
		return nil, nil, ok
	}

	ft := s.t.fields[i]

	return &ft, unsafe.Add(s.ptr, ft.off), ok
}

func (d *Decoder) decodeStructCompound(s *struct_t) error {
	var name string
	for {
		tag, err := d.readByte2()
		if err != nil {
			return err
		}

		if tag == end_t {
			return nil
		}
		if err := d.readStringNonCopy(&name); err != nil {
			return err
		}

		field, fieldPtr, ok := s.field(name)

		if ok {
			//if err := match(name, tag, field.typ); err != nil {
			//	return err
			//}

			switch tag {
			case byte_t:
				if err := d.readBytePtr(fieldPtr); err != nil {
					return err
				}
			case string_t:
				if err := d.readStringPtr(fieldPtr); err != nil {
					return err
				}
			case short_t:
				if err := d.readShortPtr(fieldPtr); err != nil {
					return err
				}
			case int_t, float_t:
				if err := d.readIntPtr(fieldPtr); err != nil {
					return err
				}
			case long_t, double_t:
				if err := d.readLongPtr(fieldPtr); err != nil {
					return err
				}
			case byte_array_t:
				switch kindOf(field.typ) {
				case array_k:
					if err := d.readByteArray(fieldPtr); err != nil {
						return err
					}
				case slice_k:
					l, err := d.readInt()
					if err != nil {
						return err
					}

					ptr := mallocgc(uintptr(l), nil, true)
					if err := d.decodeByteArray(int(l), ptr); err != nil {
						return err
					}
					setSliceArray(fieldPtr, ptr, int(l))
				}
			case long_array_t:
				switch kindOf(field.typ) {
				case array_k:
					if err := d.readLongArray(fieldPtr); err != nil {
						return err
					}
				case slice_k:
					l, err := d.readInt()
					if err != nil {
						return err
					}

					ptr := mallocgc(uintptr(l)*8, nil, true)
					if err := d.decodeLongArray(int(l), ptr); err != nil {
						return err
					}
					setSliceArray(fieldPtr, ptr, int(l))
				}
			case int_array_t:
				switch kindOf(field.typ) {
				case array_k:
					if err := d.readIntArray(fieldPtr); err != nil {
						return err
					}
				case slice_k:
					l, err := d.readInt()
					if err != nil {
						return err
					}

					ptr := mallocgc(uintptr(l)*4, nil, true)
					if err := d.decodeIntArray(int(l), ptr); err != nil {
						return err
					}
					setSliceArray(fieldPtr, ptr, int(l))
				}
			case compound_t:
				switch kindOf(field.typ) {
				case struct_k:
					s := newStruct((*structType)(unsafe.Pointer(field.typ)), fieldPtr)
					if err := d.decodeStructCompound(s); err != nil {
						return err
					}
					structs.Put(s)
				case map_k:
					t := (*mapType)(unsafe.Pointer(field.typ))
					if kK, kE := kindOf(t.key), kindOf(t.elem); kK != string_k || kE != string_k {
						return MismatchError{
							name:         name,
							expectedKind: "struct/map[string]string",
							gotKind:      fmt.Sprintf("map[%s]%s", kind_name(kK), kind_name(kE)),
							tag:          tagName(string([]byte{tag})),
						}
					}

					m := *(*map[string]string)(fieldPtr)

					if m == nil {
						m = make(map[string]string)
						*(*map[string]string)(fieldPtr) = m
					}

					if err := d.decodeMapString(m); err != nil {
						return err
					}
				}
			case list_t:
				switch kindOf(field.typ) {
				case array_k:
					t := (*arrayType)(unsafe.Pointer(field.typ))
					if err := d.decodeListArray(&array_t{t.elem, fieldPtr, t.len}); err != nil {
						return err
					}
				case slice_k:
					if err := d.decodeListSlice(fieldPtr, (*ptrType)(unsafe.Pointer(field.typ))); err != nil {
						return err
					}
				}
			default:
				return ErrUnknownTag
			}
		} else {
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
		//fmt.Println(tag, name)
	}
}

func match(name, tag_b string, t *unsafe2.Type) error {
	k := kindOf(t)
	switch tag_b {
	case byte_t_b:
		if k != uint8_k && k != int8_k && k != bool_k {
			return MismatchError{
				name:         name,
				expectedKind: "uint8/int8/bool",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case short_t_b:
		if k != int16_k && k != uint16_k {
			return MismatchError{
				name:         name,
				expectedKind: "int16/uint16",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case int_t_b:
		if k != int32_k && k != uint32_k {
			return MismatchError{
				name:         name,
				expectedKind: "int32/uint32",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case long_t_b:
		if k != int64_k && k != uint64_k {
			return MismatchError{
				name:         name,
				expectedKind: "int64/uint64",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case float_t_b:
		if k != float32_k {
			return MismatchError{
				name:         name,
				expectedKind: "float32",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case double_t_b:
		if k != float64_k {
			return MismatchError{
				name:         name,
				expectedKind: "float64",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case byte_array_t_b:
		if k != slice_k && k != array_k {
			return MismatchError{
				name:         name,
				expectedKind: "[]uint8/[]int8",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
		base := (*ptrType)(unsafe.Pointer(t))

		elemKind := kindOf(base.elem)

		if elemKind != int8_k && elemKind != uint8_k {
			return MismatchError{
				name:         name,
				expectedKind: "[]uint8/[]int8",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case string_t_b:
		if k != string_k {
			return MismatchError{
				name:         name,
				expectedKind: "string",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case list_t_b:
		if k != slice_k && k != array_k {
			return MismatchError{
				name:         name,
				expectedKind: "slice/array",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case compound_t_b:
		if k != struct_k && k != map_k {
			return MismatchError{
				name:         name,
				expectedKind: "struct/map[string]string",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case int_array_t_b:
		if k != slice_k && k != array_k {
			return MismatchError{
				name:         name,
				expectedKind: "[]uint32/[]int32",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
		base := (*ptrType)(unsafe.Pointer(t))

		elemKind := kindOf(base.elem)

		if elemKind != int32_k && elemKind != uint32_k {
			return MismatchError{
				name:         name,
				expectedKind: "[]uint32/[]int32",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	case long_array_t_b:
		if k != slice_k && k != array_k {
			return MismatchError{
				name:         name,
				expectedKind: "[]uint64/[]int64",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
		base := (*ptrType)(unsafe.Pointer(t))

		elemKind := kindOf(base.elem)

		if elemKind != int64_k && elemKind != uint64_k {
			return MismatchError{
				name:         name,
				expectedKind: "[]uint64/[]int64",
				gotKind:      kind_name(k),
				tag:          tagName(tag_b),
			}
		}
	}
	return nil
}
