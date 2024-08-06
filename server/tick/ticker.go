package tick

import "sync/atomic"

// New creates a new tick manager with freq ticks per second
func New(freq int) *TickManager {
	mgr := &TickManager{}
	mgr.tickingFrequency.Store(int32(freq))

	return mgr
}

type TickManager struct {
	tickingFrequency atomic.Int32
}

func (mgr *TickManager) SetFrequency() {

}
