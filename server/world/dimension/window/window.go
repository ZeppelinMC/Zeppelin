package window

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
)

var windowId atomic.Int32

type WindowManager struct {
	windows map[pos.BlockPosition]*Window
	mu      sync.RWMutex
}

func NewManager() *WindowManager {
	return &WindowManager{windows: make(map[pos.BlockPosition]*Window)}
}

func (mgr *WindowManager) AddWindow(position pos.BlockPosition, w *Window) error {
	if w.Id == 0 {
		return fmt.Errorf("window id should not be 0")
	}
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.windows[position] = w
	return nil
}

func (mgr *WindowManager) Range(itr func(pos.BlockPosition, *Window)) {
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

func (mgr *WindowManager) Get(id int32) (pos.BlockPosition, *Window, bool) {
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
func (mgr *WindowManager) New(windowType, chunkEntityType string, items container.Container, title text.TextComponent) *Window {
	return &Window{
		mgr:             mgr,
		Items:           items,
		Id:              windowId.Add(1),
		Title:           title,
		Type:            windowType,
		ChunkEntityType: chunkEntityType,
	}
}

type Window struct {
	mgr *WindowManager

	ChunkEntityType string

	Items container.Container
	Id    int32
	Type  string
	Title text.TextComponent

	Viewers byte
}

func (w *Window) Position() (x, y, z int32, ok bool) {
	pos, _, ok := w.mgr.Get(w.Id)

	return pos[0], pos[1], pos[2], ok
}
