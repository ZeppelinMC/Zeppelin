package entity

import (
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type server interface {
	SetEntityPosition(id int32, x, y, z float64)
	TeleportEntity(id int32, x, y, z float64)
}

type Entity interface {
	// Happens every tick
	Tick(srv server, tick uint)
}

var pool = map[string]func(int32, chunk.Entity, *world.Dimension) Entity{}

func NewEntity(data chunk.Entity, id int32, d *world.Dimension) Entity {
	if f, ok := pool[data.Id]; ok {
		return f(id, data, d)
	} else {
		return &UnknownEntity{entityId: id, data: data, dimension: d}
	}
}

type UnknownEntity struct {
	entityId  int32
	data      chunk.Entity
	dimension *world.Dimension
}

func (e *UnknownEntity) Tick(srv server, tick uint) {
	x, y, z := e.data.Pos[0], e.data.Pos[1], e.data.Pos[2]
	if b := e.dimension.Block(int64(x), int64(y-1), int64(z)); b == nil || b.EncodedName() == "minecraft:air" {
		y--
		srv.SetEntityPosition(e.entityId, x, y, z)
	}

}
