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
	uuid, _ := world.IntUUIDToByteUUID(data.UUID)
	id := idCounter.Add(1)
	e := &Entity{data, id, uuid}
	srv.entities[id] = e
	return e
}
