package qnbt

import (
	"unsafe"

	"github.com/oq-x/unsafe2"
)

//go:linkname mallocgc runtime.mallocgc
func mallocgc(size uintptr, typ *unsafe2.Type, needzero bool) unsafe.Pointer
