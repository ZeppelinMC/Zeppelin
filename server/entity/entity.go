package entity

import (
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
	Type() string

	Position() (x, y, z float64)
	Rotation() (yaw, pitch float32)
	OnGround() bool
	SetPosition(x, y, z float64)
}

type LivingEntity interface {
	Entity
	Kill()
	Attack(attacker LivingEntity)
	Health() float32
	SetHealth(f float32)
}

type newf func(int32, uuid.UUID, chunk.Entity, *world.Dimension) Entity

var pool = map[string]newf{}

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

func Register(typ string, f newf) {
	pool[typ] = f
}
