package qnbt

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/oq-x/unsafe2"
)

func (d *Decoder) decodeListSlice(ptr unsafe.Pointer, t *ptrType) error {
	tag, err := d.readByte2()
	if err != nil {
		return err
	}

	l, err := d.readInt()
	if err != nil {
		return err
	}

	if tag == end_t {
		return nil
	}

	cap := *(*int)(unsafe.Add(ptr, ptrSize*2))
	array := &array_t{
		elem: t.elem,
		ptr:  unsafe.Pointer(*(*uintptr)(ptr)),
		len:  uintptr(cap),
	}

	if int(l) > cap {
		array.ptr = mallocgc(uintptr(l)*array.elem.Size, nil, true)
		array.len = uintptr(l)
		setSliceArray(ptr, array.ptr, int(l))
	}

	return d.readListArray(tag, l, array)
}

type array_t struct {
	elem *unsafe2.Type
	ptr  unsafe.Pointer
	len  uintptr
}

func (a *array_t) index(i int) unsafe.Pointer {
	return unsafe.Add(a.ptr, a.elem.Size*uintptr(i))
}

func (d *Decoder) decodeListArray(a *array_t) error {
	tag, err := d.readByte2()
	if err != nil {
		return err
	}

	l, err := d.readInt()
	if err != nil {
		return err
	}

	if tag == end_t {
		return nil
	}

	if uintptr(l) < a.len {
		return fmt.Errorf("array too small")
	}

	return d.readListArray(tag, l, a)
}

func (d *Decoder) readListArray(tag byte, l int32, a *array_t) error {
	kind := kindOf(a.elem)

	for i := 0; i < int(l); i++ {
		iptr := a.index(i)

		//if err := match("", tag, a.elem); err != nil {
		//	return err
		//}

		switch tag {
		case byte_t:
			if err := d.readBytePtr(iptr); err != nil {
				return err
			}
		case short_t:
			if err := d.readShortPtr(iptr); err != nil {
				return err
			}
		case int_t, float_t:
			if err := d.readIntPtr(iptr); err != nil {
				return err
			}
		case long_t, double_t:
			if err := d.readLongPtr(iptr); err != nil {
				return err
			}
		case byte_array_t:
			switch kind {
			case array_k:
				if err := d.readByteArray(iptr); err != nil {
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
				setSliceArray(iptr, ptr, int(l))
			}
		case string_t:
			if err := d.readStringPtr(iptr); err != nil {
				return err
			}
		case long_array_t:
			switch kind {
			case array_k:
				if err := d.readLongArray(iptr); err != nil {
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
				setSliceArray(iptr, ptr, int(l))
			}
		case int_array_t:
			switch kind {
			case array_k:
				if err := d.readIntArray(iptr); err != nil {
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
				setSliceArray(iptr, ptr, int(l))
			}
		case compound_t:
			switch kind {
			case struct_k:
				s := newStruct((*structType)(unsafe.Pointer(a.elem)), iptr)
				if err := d.decodeStructCompound(s); err != nil {
					return err
				}
				structs.Put(s)
			case map_k:
				t := (*mapType)(unsafe.Pointer(a.elem))
				if kK, kE := kindOf(t.key), kindOf(t.elem); kK != string_k || kE != string_k {
					return MismatchError{
						name:         strconv.Itoa(i),
						expectedKind: "struct/map[string]string",
						gotKind:      fmt.Sprintf("map[%s]%s", kind_name(kK), kind_name(kE)),
						//tag:          tagName(tag),
					}
				}

				m := *(*map[string]string)(iptr)

				if m == nil {
					m = make(map[string]string)
					*(*map[string]string)(iptr) = m
				}

				if err := d.decodeMapString(m); err != nil {
					return err
				}
			}
		case list_t:
			switch kind {
			case array_k:
				arr := (*arrayType)(unsafe.Pointer(a.elem))
				if err := d.decodeListArray(&array_t{arr.elem, iptr, arr.len}); err != nil {
					return err
				}
			case slice_k:
				if err := d.decodeListSlice(iptr, (*ptrType)(unsafe.Pointer(a.elem))); err != nil {
					return err
				}
			}
		default:
			return ErrUnknownTag
		}
	}
	return nil
}

func (d *Decoder) discardList() error {
	tag, err := d.readByte2()
	if err != nil {
		return err
	}

	len, err := d.readInt()
	if err != nil {
		return err
	}

	f := fixedSize(tag)
	if f == 0 {
		return nil
	}
	if f != -1 {
		return d.rd.skip(f * int(len))
	}

	for i := int32(0); i < len; i++ {
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

	return nil
}

func fixedSize(tag byte) int {
	switch tag {
	case end_t:
		return 0
	case byte_t:
		return 1
	case short_t:
		return 2
	case int_t, float_t:
		return 4
	case long_t, double_t:
		return 8
	default:
		return -1
	}
}
