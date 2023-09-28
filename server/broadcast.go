package server

import (
	"fmt"
	"math"

	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/player"
	"github.com/dynamitemc/dynamite/server/commands"
)

func (srv *Server) GlobalBroadcast(pk packet.Packet) {
	for _, p := range srv.Players {
		p.session.SendPacket(pk)
	}
}

func (srv *Server) GlobalMessage(message string) {
	srv.GlobalBroadcast(&packet.SystemChatMessage{
		Content: message,
	})
	fmt.Println(commands.ParseChat(message))
}

func (p *PlayerController) BroadcastMovement(oldx, oldy, oldz float64) {
	x1, y1, z1 := p.Position()
	ong := p.OnGround()
	for _, pl := range p.Server.Players {
		if pl.UUID == p.UUID {
			continue
		}
		x2, y2, z2 := pl.Position()
		distance := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1))
		if pl.ClientSettings().ViewDistance*16 > int8(distance) {
			continue
		}
		pl.session.SendPacket(&packet.EntityPosition{
			EntityID: p.player.EntityID,
			X:        (int16(x1)*32 - int16(oldx)*32) * 128,
			Y:        (int16(y1)*32 - int16(oldy)*32) * 128,
			Z:        (int16(z1)*32 - int16(oldz)*32) * 128,
			OnGround: ong,
		})
	}
}

func (p *PlayerController) Spawn() {
	x, y, z := p.Position()
	yaw, pitch := p.Rotation()
	for _, pl := range p.Server.Players {
		if pl.UUID == p.UUID {
			continue
		}
		fmt.Println(p.player.EntityID, pl.player.EntityID)
		pl.session.SendPacket(&packet.SpawnPlayer{
			EntityID:   p.player.EntityID,
			PlayerUUID: p.session.Info().UUID,
			X:          x,
			Y:          y,
			Z:          z,
			Yaw:        byte(yaw),
			Pitch:      byte(pitch),
		})
	}
}

func (srv *Server) PlayerlistUpdate() {
	var players []player.Info
	for _, p := range srv.Players {
		p.session.Info().Listed = true
		players = append(players, *p.session.Info())
	}
	srv.GlobalBroadcast(&packet.PlayerInfoUpdate{
		Actions: 0x01 | 0x08,
		Players: players,
	})
}

func (srv *Server) PlayerlistRemove(players ...[16]byte) {
	srv.GlobalBroadcast(&packet.PlayerInfoRemove{UUIDS: players})
}
