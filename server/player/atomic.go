package player

import (
	"sync/atomic"
	"unsafe"
)

func u64f(i uint64) float64 {
	return *(*float64)(unsafe.Pointer(&i))
}
func f64u(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func u32f(i uint32) float32 {
	return *(*float32)(unsafe.Pointer(&i))
}
func f32u(f float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&f))
}

func atomicFloat64(f float64) atomic.Uint64 {
	var v atomic.Uint64
	v.Store(f64u(f))

	return v
}
func atomicFloat32(f float32) atomic.Uint32 {
	return atomicUint32(f32u(f))
}

func atomicUint32(i uint32) atomic.Uint32 {
	var v atomic.Uint32
	v.Store(i)

	return v
}

func atomicInt32(i int32) atomic.Int32 {
	var v atomic.Int32
	v.Store(i)

	return v
}

func atomicBool(b bool) atomic.Bool {
	var v atomic.Bool
	v.Store(b)

	return v
}
