package player

import (
	"sync/atomic"
	"unsafe"
)

func i64f(i int64) float64 {
	return *(*float64)(unsafe.Pointer(&i))
}
func f64i(f float64) int64 {
	return *(*int64)(unsafe.Pointer(&f))
}

func i32f(i int32) float32 {
	return *(*float32)(unsafe.Pointer(&i))
}
func f32i(f float32) int32 {
	return *(*int32)(unsafe.Pointer(&f))
}

func atomicFloat64(f float64) *atomic.Int64 {
	var v *atomic.Int64
	v.Store(f64i(f))

	return v
}
func atomicFloat32(f float32) *atomic.Int32 {
	return atomicInt32(f32i(f))
}

func atomicInt32(i int32) *atomic.Int32 {
	var v *atomic.Int32
	v.Store(i)

	return v
}

func atomicBool(b bool) *atomic.Bool {
	var v *atomic.Bool
	v.Store(b)

	return v
}
