package server

import (
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type Entity struct {
	data chunk.Entity
	ID   int32
	UUID [16]byte
}

func (srv *Server) NewEntity(data chunk.Entity) *Entity {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	uuid, _ := world.NBTToUUID(data.UUID)
	id := srv.entityCounter.Add(1)
	e := &Entity{data, id, uuid}
	srv.Entities[id] = e
	return e
}
