package server

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/entity"
	"github.com/dynamitemc/dynamite/server/registry"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
	"github.com/google/uuid"
)

type Entity struct {
	data   chunk.Entity
	Entity entity.Entity
	ID     int32
	UUID   [16]byte
}

func (srv *Server) NewEntity(data chunk.Entity, d *world.Dimension) *Entity {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	uuid, _ := world.IntUUIDToByteUUID(data.UUID)
	id := idCounter.Add(1)
	e := &Entity{data: data, ID: id, UUID: uuid, Entity: entity.NewEntity(data, id, d)}
	srv.entities[id] = e
	return e
}

func (srv *Server) FindEntity(id int32) interface{} {
	if p := srv.FindPlayerByID(id); p != nil {
		return p
	} else {
		srv.mu.RLock()
		defer srv.mu.RUnlock()
		e, ok := srv.entities[id]
		if !ok {
			return nil
		}
		return e
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

func (srv *Server) SpawnEntity(typ string, x, y, z float64) {
	id, ok := registry.GetEntity(typ)
	if !ok {
		return
	}
	e := srv.NewEntity(chunk.Entity{
		Pos:  []float64{x, y, z},
		Id:   typ,
		UUID: world.ByteUUIDToIntUUID(uuid.New()),
	}, srv.World.Overworld())
	p := &packet.SpawnEntity{
		EntityID: e.ID,
		UUID:     e.UUID,
		X:        x,
		Y:        y,
		Z:        z,
		Type:     id.ProtocolID,
	}
	srv.mu.RLock()
	defer srv.mu.RUnlock()

	for _, pl := range srv.players {
		if !pl.InView2(x, y, z) {
			continue
		}
		pl.mu.Lock()
		pl.spawnedEntities = append(pl.spawnedEntities, e.ID)
		pl.mu.Unlock()
		pl.SendPacket(p)
	}
}

func (srv *Server) SetEntityPosition(id int32, x, y, z float64) {
	entity := srv.FindEntity(id)
	if entity == nil {
		return
	}
	if e, ok := entity.(*Entity); ok {
		prevX, prevY, prevZ := e.data.Pos[0], e.data.Pos[1], e.data.Pos[2]

		// todo check distance
		dx := int16((x*32 - prevX*32) * 128)
		dy := int16((y*32 - prevY*32) * 128)
		dz := int16((z*32 - prevZ*32) * 128)

		e.data.Pos[0], e.data.Pos[1], e.data.Pos[2] = x, y, z
		ong := true
		if e.data.OnGround == 0 {
			ong = false
		}

		for _, pl := range srv.players {
			if !pl.IsSpawned(id) || !pl.InView2(x, y, z) {
				continue
			}
			pl.SendPacket(&packet.EntityPosition{
				EntityID: id,
				OnGround: ong,
				X:        dx,
				Y:        dy,
				Z:        dz,
			})
		}
	}
}

func (srv *Server) TeleportEntity(id int32, x, y, z float64) {
	entity := srv.FindEntity(id)
	if entity == nil {
		return
	}
	if e, ok := entity.(*Entity); ok {
		e.data.Pos[0], e.data.Pos[1], e.data.Pos[2] = x, y, z
		srv.mu.RLock()
		defer srv.mu.RUnlock()
		ong := true
		if e.data.OnGround == 0 {
			ong = false
		}

		for _, pl := range srv.players {
			if !pl.IsSpawned(id) || !pl.InView2(x, y, z) {
				continue
			}
			pl.SendPacket(&packet.TeleportEntity{
				EntityID: id,
				OnGround: ong,
				X:        x,
				Y:        y,
				Z:        z,
			})
		}
	}
}
