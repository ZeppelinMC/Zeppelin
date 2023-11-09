package controller

import "sync"

type Controller[K comparable, V any] struct {
	mu      sync.RWMutex
	entries map[K]V
}

// New creates a new controller
func New[K comparable, V any]() *Controller[K, V] {
	return &Controller[K, V]{
		entries: make(map[K]V),
	}
}

// Count returns the amount of entries in the controller
func (controller *Controller[K, V]) Count() int {
	controller.mu.RLock()
	defer controller.mu.RUnlock()
	return len(controller.entries)
}

// Range iterates over the entries and locks the entry map
// use return true in an iteration to continue and return false to break it
// returns the amount of entries iterated and the entry that broke the iteration if exists
func (controller *Controller[K, V]) Range(f func(K, V) bool) (int, V) {
	var iterated int
	var val V
	controller.mu.RLock()
	defer controller.mu.RUnlock()
	for k, v := range controller.entries {
		iterated++
		if !f(k, v) {
			val = v
			break
		}
	}
	return iterated, val
}

// Works the same way as Range but doesn't lock the entry map
func (controller *Controller[K, V]) RangeNoLock(f func(K, V) bool) (int, V) {
	var iterated int
	var val V
	for k, v := range controller.entries {
		iterated++
		if !f(k, v) {
			val = v
			break
		}
	}
	return iterated, val
}

func (controller *Controller[K, V]) Set(k K, v V) {
	controller.mu.Lock()
	defer controller.mu.Unlock()
	controller.entries[k] = v
}

func (controller *Controller[K, V]) Get(k K) V {
	controller.mu.RLock()
	defer controller.mu.RUnlock()
	return controller.entries[k]
}

func (controller *Controller[K, V]) Get2(k K) (v V, ok bool) {
	controller.mu.RLock()
	defer controller.mu.RUnlock()
	v, ok = controller.entries[k]
	return
}

func (controller *Controller[K, V]) Find(find func(k K, v V) bool) (v V) {
	_, d := controller.Range(func(k K, v V) bool {
		return !find(k, v)
	})
	return d
}

func (controller *Controller[K, V]) Delete(k K) {
	controller.mu.Lock()
	defer controller.mu.Unlock()
	delete(controller.entries, k)
}
