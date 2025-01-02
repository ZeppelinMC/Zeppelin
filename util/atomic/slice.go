package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Pointer is the size of pointers
var Pointer = unsafe.Sizeof(0)

//go:linkname mallocgc runtime.mallocgc
func mallocgc(size uintptr, typ unsafe.Pointer, needzero bool) unsafe.Pointer

func memcpy(dst, src unsafe.Pointer, n uintptr) {
	dstS := unsafe.Slice((*byte)(dst), n)
	srcS := unsafe.Slice((*byte)(src), n)

	copy(dstS, srcS)
}

func Make(len, cap, sizeOfElement uintptr) Slice {
	return Slice{
		ptr: uintptr(mallocgc(cap, nil, false)),
		len: len, cap: cap,

		sizeOfElement: sizeOfElement,
	}
}

// Slice is an atomic slice with all read only fields. Entries cannot be removed.
type Slice struct {
	ptr, len, cap uintptr

	//not atomic
	sizeOfElement uintptr
}

func (a *Slice) Len() uintptr {
	return atomic.LoadUintptr(&a.len)
}

func (a *Slice) Cap() uintptr {
	return atomic.LoadUintptr(&a.cap)
}

func (a *Slice) Pointer() unsafe.Pointer {
	return unsafe.Pointer(atomic.LoadUintptr(&a.ptr))
}

// At returns the element at the specified index, while checking that the index is in bounds
func (a *Slice) At(index uintptr) unsafe.Pointer {
	if a.Cap() < index {
		panic("out of bounds")
	}
	if a.Len() < index {
		atomic.StoreUintptr(&a.len, index+1)
	}
	return a.Element(index)
}

// Element returns the element at the specified index, without checking that the index is in bounds
func (a *Slice) Element(index uintptr) unsafe.Pointer {
	return unsafe.Add(a.Pointer(), index*a.sizeOfElement)
}

// Grow grows the slice to capacity of cap
func (a *Slice) Grow(cap uintptr) {
	newPtr := mallocgc(cap*a.sizeOfElement, nil, false)
	length := a.Len()

	memcpy(newPtr, a.Pointer(), a.sizeOfElement*length)

	atomic.StoreUintptr(&a.ptr, uintptr(newPtr))
	atomic.StoreUintptr(&a.cap, cap)
}

// TrimStart removes the first count elements from the slice
func (a *Slice) TrimStart(count uintptr) {
	if count > a.Cap() {
		panic("out of bounds")
	}

	atomic.AddUintptr(&a.ptr, count*a.sizeOfElement)
}

// TrimEnd removes the last count elements from the slice
func (a *Slice) TrimEnd(count uintptr) {
	if count > a.Cap() {
		panic("out of bounds")
	}

	atomic.AddUintptr(&a.len, -count*a.sizeOfElement)
}

// Append adds the elements from the specified pointers to the slice
func (a *Slice) Append(elements ...unsafe.Pointer) {
	currentLength := a.Len()
	currentCapacity := a.Cap()

	count := uintptr(len(elements))

	// need to expand the slice
	if newLength := currentLength + count; newLength > currentCapacity {
		a.Grow(newLength)
		a.Append(elements...)
		return
	}

	basePtr := unsafe.Add(a.Pointer(), currentLength*a.sizeOfElement)

	for _, element := range elements {
		memcpy(basePtr, element, a.sizeOfElement)
		basePtr = unsafe.Add(basePtr, a.sizeOfElement)
	}

	atomic.AddUintptr(&a.len, count)
}
