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

func (srv *Server) FindEntity(id int32) interface{} {
	if p := srv.FindPlayerByID(id); p != nil {
		return p
	} else {
		srv.mu.RLock()
		defer srv.mu.RUnlock()
		return srv.entities[id]
	}
}

func (srv *Server) FindEntityByUUID(id [16]byte) interface{} {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.players {
		if p.conn.UUID() == id {
			return p
		}
	}
	for _, e := range srv.entities {
		if e.UUID == id {
			return e
		}
	}
	return nil
}
