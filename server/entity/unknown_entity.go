package entity

import (
	"sync"

	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
	"github.com/google/uuid"
)

type UnknownEntity struct {
	mu        sync.RWMutex
	entityId  int32
	uuid      uuid.UUID
	data      chunk.Entity
	dimension *world.Dimension
}

func (e *UnknownEntity) Type() string {
	return e.data.Id
}

func (e *UnknownEntity) UUID() uuid.UUID {
	return e.uuid
}

func (e *UnknownEntity) EntityID() int32 {
	return e.entityId
}

func (e *UnknownEntity) Position() (x, y, z float64) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.data.Pos[0], e.data.Pos[1], e.data.Pos[2]
}

func (e *UnknownEntity) SetPosition(x, y, z float64) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.data.Pos[0], e.data.Pos[1], e.data.Pos[2] = x, y, z
}

func (e *UnknownEntity) Rotation() (yaw, pitch float32) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.data.Rotation[0], e.data.Rotation[1]
}

func (e *UnknownEntity) OnGround() (ong bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	ong = true
	if e.data.OnGround == 0 {
		ong = false
	}
	return
}

func (e *UnknownEntity) Tick(srv server, tick uint) {
	x, y, z := e.data.Pos[0], e.data.Pos[1], e.data.Pos[2]
	if b := e.dimension.Block(int64(x), int64(y-1), int64(z)); b == nil || b.EncodedName() == "minecraft:air" {
		y--
		srv.SetEntityPosition(e.entityId, x, y, z)
	}

}
