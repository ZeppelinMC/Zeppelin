package state

import (
	"sync"

	"github.com/google/uuid"
)

// the player cache is used when saving the world playerdata

type PlayerEntityManager struct {
	m  map[uuid.UUID]*PlayerEntity
	mu sync.RWMutex
}

func (mgr *PlayerEntityManager) lookup(id uuid.UUID) (*PlayerEntity, bool) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	p, ok := mgr.m[id]

	return p, ok
}

func (mgr *PlayerEntityManager) add(p *PlayerEntity) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.m[p.UUID()] = p
}

// saves all the players in the manager to the path specified in their data file
func (mgr *PlayerEntityManager) SaveAll() {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	for _, player := range mgr.m {
		player.sync()

		player.data.Save()
	}
}

func NewPlayerEntityManager() *PlayerEntityManager {
	return &PlayerEntityManager{m: make(map[uuid.UUID]*PlayerEntity)}
}
