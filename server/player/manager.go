package player

import (
	"sync"

	"github.com/google/uuid"
)

// the player cache is used when saving the world playerdata

type PlayerManager struct {
	m  map[uuid.UUID]*Player
	mu sync.RWMutex
}

func (cache *PlayerManager) lookup(id uuid.UUID) (*Player, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	p, ok := cache.m[id]

	return p, ok
}

func (cache *PlayerManager) add(p *Player) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.m[p.UUID()] = p
}

// saves all the players in the manager to the path specified in their data file
func (cache *PlayerManager) SaveAll() {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	for _, player := range cache.m {
		player.sync()

		player.data.Save()
	}
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{m: make(map[uuid.UUID]*Player)}
}
