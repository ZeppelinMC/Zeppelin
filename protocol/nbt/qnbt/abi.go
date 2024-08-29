package qnbt

import (
	"unsafe"

	"github.com/oq-x/unsafe2"
)

type name struct {
	Bytes *byte
}

func (n name) exported() bool {
	return (*n.Bytes)&(1<<0) != 0
}

func (n name) hastag() bool {
	return (*n.Bytes)&(1<<1) != 0
}

func (n name) embedded() bool {
	return (*n.Bytes)&(1<<3) != 0
}

func (n name) rvi(off int) (int, int) {
	v := 0
	for i := 0; ; i++ {
		x := *n.data(off + i)
		v += int(x&0x7f) << (7 * i)
		if x&0x80 == 0 {
			return i + 1, v
		}
	}
}

func (n name) data(off int) *byte {
	return (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(n.Bytes)) + uintptr(off)))
}

func (n name) name() string {
	i, l := n.rvi(1)
	return unsafe.String(n.data(i+1), l)
}

func (n name) tag() string {
	if !n.hastag() {
		return ""
	}
	i, l := n.rvi(1)
	i2, l2 := n.rvi(1 + i + l)
	return unsafe.String(n.data(1+i+l+i2), l2)
}

type structField struct {
	name name
	typ  *unsafe2.Type
	off  uintptr
}

type structType struct {
	unsafe2.Type
	pkgPath name
	fields  []structField
}

type ptrType struct {
	unsafe2.Type
	elem *unsafe2.Type
}

type arrayType struct {
	unsafe2.Type
	elem  *unsafe2.Type
	slice *unsafe2.Type
	len   uintptr
}

type mapType struct {
	unsafe2.Type
	key  *unsafe2.Type
	elem *unsafe2.Type
	// rest is unimportant
}

// returns the element of a pointer
func ptrelem(t *unsafe2.Type, v uintptr) (*unsafe2.Type, unsafe.Pointer) {
	ptr_t := *(*ptrType)(unsafe.Pointer(t))

	elem_t := ptr_t.elem

	return elem_t, unsafe.Pointer(v)
}

func kindOf(t *unsafe2.Type) uint8 {
	return t.Kind & kindm
}

const (
	kindm = (1 << 5) - 1
)

func kind_name(k uint8) string {
	switch k {
	case bool_k:
		return "bool"

	case int_k:
		return "int"
	case int8_k:
		return "int8"
	case int16_k:
		return "int16"
	case int32_k:
		return "int32"
	case int64_k:
		return "int64"

	case uint_k:
		return "uint"
	case uint8_k:
		return "uint8"
	case uint16_k:
		return "uint16"
	case uint32_k:
		return "uint32"
	case uint64_k:
		return "uint64"
	case uintptr_k:
		return "uintptr"

	case float32_k:
		return "float32"
	case float64_k:
		return "float64"
	case complex64_k:
		return "complex64"
	case complex128_k:
		return "complex128"

	case array_k:
		return "array"
	case chan_k:
		return "chan"
	case func_k:
		return "func"
	case interface_k:
		return "interface"
	case map_k:
		return "map"
	case pointer_k:
		return "pointer"
	case slice_k:
		return "slice"
	case string_k:
		return "string"
	case struct_k:
		return "struct"
	case unsafe_ptr_k:
		return "unsafe.Pointer"
	default:
		return "invalid"
	}
}

const (
	invalid_k uint8 = iota
	bool_k
	int_k
	int8_k
	int16_k
	int32_k
	int64_k
	uint_k
	uint8_k
	uint16_k
	uint32_k
	uint64_k
	uintptr_k
	float32_k
	float64_k
	complex64_k
	complex128_k
	array_k
	chan_k
	func_k
	interface_k
	map_k
	pointer_k
	slice_k
	string_k
	struct_k
	unsafe_ptr_k
)

func setSliceArray(s, a unsafe.Pointer, l int) {
	*(*uintptr)(s) = uintptr(a)
	*(*int)(unsafe.Add(s, ptrSize)) = l
	*(*int)(unsafe.Add(s, ptrSize*2)) = l
}

var ptrSize = unsafe.Sizeof(0)
