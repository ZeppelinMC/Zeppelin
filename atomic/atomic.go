package atomic

import "sync/atomic"

type AtomicValue[T any] struct {
	v atomic.Value
}

func (a *AtomicValue[T]) Get() T {
	val := a.v.Load()
	if val == nil {
		var e T
		return e
	}
	v, _ := val.(T)

	return v
}

func (a *AtomicValue[T]) Set(t T) {
	a.v.Store(t)
}

func Value[T any](value T) AtomicValue[T] {
	val := AtomicValue[T]{}
	val.Set(value)
	return val
}
