package atomic

import "sync/atomic"

type AtomicValue[T any] struct {
	v atomic.Value
}

func (a *AtomicValue[T]) Get() T {
	if a.v.Load() == nil {
		var e T
		return e
	}
	return a.v.Load().(T)
}

func (a *AtomicValue[T]) Set(t T) {
	a.v.Store(t)
}
