package server

import (
	"time"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/entity"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/google/uuid"
)

func (srv *Server) tickLoop() {
	var n uint
	for range time.Tick(time.Second / time.Duration(srv.Config.TPS)) {
		srv.tick(n)
		n++
	}
}

func (srv *Server) tick(tick uint) {
	srv.Entities.RangeNoLock(func(_ int32, e entity.Entity) bool {
		e.Tick(srv, tick)
		return true
	})

	worldAge, dayTime := srv.World.IncrementTime()
	srv.Players.RangeNoLock(func(_ uuid.UUID, pl *player.Player) bool {
		if tick%8 == 0 {
			pl.SendChunks(pl.Dimension())
			//pl.UnloadChunks()
		}

		pl.SendPacket(&packet.UpdateTime{
			WorldAge:  worldAge,
			TimeOfDay: dayTime,
		})
		return true
	})
}
