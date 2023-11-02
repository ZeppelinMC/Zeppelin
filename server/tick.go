package server

import (
	"time"

	"github.com/aimjel/minecraft/packet"
)

func (srv *Server) tickLoop() {
	var n uint
	for range time.Tick(time.Second / time.Duration(srv.Config.TPS)) {
		//srv.tick(n)
		n++
	}
}

func (srv *Server) tick(tick uint) {
	for _, pl := range srv.players {
		//if tick%8 == 0 {
		//pl.SendChunks(srv.GetDimension(pl.Player.Dimension()))
		//pl.UnloadChunks()
		//}

		worldAge, dayTime := srv.World.IncrementTime()
		pl.SendPacket(&packet.UpdateTime{
			WorldAge:  worldAge,
			TimeOfDay: dayTime,
		})
	}
}
