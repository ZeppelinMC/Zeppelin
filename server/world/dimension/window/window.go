package window

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/text"
)

var windowId atomic.Int32

type WindowManager struct {
	windows map[[3]int32]*Window
	mu      sync.RWMutex
}

func NewManager() *WindowManager {
	return &WindowManager{windows: make(map[[3]int32]*Window)}
}

func (mgr *WindowManager) AddWindow(position [3]int32, w *Window) error {
	if w.Id == 0 {
		return fmt.Errorf("window id should not be 0")
	}
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.windows[position] = w
	return nil
}

func (mgr *WindowManager) Range(itr func([3]int32, *Window)) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	for pos, w := range mgr.windows {
		itr(pos, w)
	}
}

func (mgr *WindowManager) At(x, y, z int32) (*Window, bool) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	w, ok := mgr.windows[[3]int32{x, y, z}]
	return w, ok
}

func (mgr *WindowManager) Get(id int32) ([3]int32, *Window, bool) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	for pos, w := range mgr.windows {
		if w.Id == id {
			return pos, w, true
		}
	}
	return [3]int32{}, nil, false
}

// window types are found at https://wiki.vg/Inventory
func New(typ string, items container.Container, title text.TextComponent) *Window {
	return &Window{
		Items: items,
		Id:    windowId.Add(1),
		Title: title,
		Type:  typ,
	}
}

type Window struct {
	Items container.Container
	Id    int32
	Type  string
	Title text.TextComponent

	Viewers byte
}
