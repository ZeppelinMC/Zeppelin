package entity

import (
	"sync"

	"github.com/dynamitemc/dynamite/server/controller"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
	"github.com/google/uuid"
)

type server interface {
	SetEntityPosition(id int32, x, y, z float64)
	TeleportEntity(id int32, x, y, z float64)
}

type Entity interface {
	Tick(srv server, tick uint)
	UUID() uuid.UUID
	EntityID() int32
	Position() (x, y, z float64)
	Rotation() (yaw, pitch float32)
	OnGround() bool
	SetPosition(x, y, z float64)
	Type() string
}

var pool = map[string]func(int32, uuid.UUID, chunk.Entity, *world.Dimension) Entity{}

func CreateEntity(entityController *controller.Controller[int32, Entity], id int32, data chunk.Entity, d *world.Dimension) Entity {
	uuid, _ := world.IntUUIDToByteUUID(data.UUID)
	e := NewEntity(data, id, uuid, d)
	entityController.Set(id, e)
	return e
}

func NewEntity(data chunk.Entity, id int32, uuid uuid.UUID, d *world.Dimension) Entity {
	if f, ok := pool[data.Id]; ok {
		return f(id, uuid, data, d)
	} else {
		return &UnknownEntity{entityId: id, data: data, dimension: d, uuid: uuid}
	}
}

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
