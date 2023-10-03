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

func (srv *Server) GlobalMessage(message string, sender *PlayerController) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if p.ClientSettings().ChatMode == 2 {
			continue
		} else if p.ClientSettings().ChatMode == 1 && sender != nil {
			continue
		}
		p.session.SendPacket(&packet.SystemChatMessage{
			Content: message,
		})
	}
	fmt.Println(commands.ParseChat(message))
}

func (p *PlayerController) BroadcastMovement(x1, y1, z1 float64, yaw, pitch float32, ong bool) {
	oldx, oldy, oldz := p.player.Position()
	p.player.SetPosition(x1, y1, z1, yaw, pitch, ong)
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.Players {
		if pl.UUID == p.UUID {
			continue
		}
		x2, y2, z2 := pl.player.Position()
		distance := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1))
		if float64(pl.ClientSettings().ViewDistance)*16 < distance {
			if pl.IsSpawned(p.player.EntityId()) {
				fmt.Println(p.Name(), "is too far! dispawning for", pl.Name())
				pl.DespawnPlayer(p)
			}
			continue
		}

		if pl.IsSpawned(p.player.EntityId()) {
			pl.session.SendPacket(&packet.EntityPositionRotation{
				EntityID: p.player.EntityId(),
				X:        ((int16(x1) * 32) - int16(oldx)*32) * 128,
				Y:        ((int16(y1) * 32) - int16(oldy)*32) * 128,
				Z:        ((int16(z1) * 32) - int16(oldz)*32) * 128,
				Yaw:      byte(yaw),
				Pitch:    byte(pitch),
				OnGround: ong,
			})
		} else {
			pl.SpawnPlayer(p)
		}
	}
}

func (p *PlayerController) Spawn() {
	x1, y1, z1 := p.player.Position()
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.Players {
		x2, y2, z2 := pl.player.Position()
		if pl.UUID == p.UUID {
			continue
		}

		distance := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1))
		if float64(pl.ClientSettings().ViewDistance)*16 >= distance {
			// notify existing players that a new player has joined
			pl.SpawnPlayer(p)
			fmt.Println("Player", p.Name(), "joined,", pl.Name(), "can see")
			fmt.Println(pl.spawnedEntities)
		} else {
			fmt.Println("Player", p.Name(), "joined,", pl.Name(), "cannot see")
		}
		if float64(p.ClientSettings().ViewDistance)*16 >= distance {
			// notify the new player of the existing players
			p.SpawnPlayer(pl)
			fmt.Println("Player", p.Name(), "joined, can see", pl.Name())
			fmt.Println(p.spawnedEntities)
		} else {
			fmt.Println("Player", p.Name(), "joined, cannot see", pl.Name())
		}
	}
}

func (srv *Server) PlayerlistUpdate() {
	var players []player.Info
	srv.mu.RLock()
	defer srv.mu.RUnlock()
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
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	srv.GlobalBroadcast(&packet.PlayerInfoRemove{UUIDS: players})
}
