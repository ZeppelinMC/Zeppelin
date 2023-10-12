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
	srv.entityCounter++
	uuid, _ := world.NBTToUUID(data.UUID)
	e := &Entity{data, srv.entityCounter, uuid}
	srv.Entities[srv.entityCounter] = e
	return e
}
