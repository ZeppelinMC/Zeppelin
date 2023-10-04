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

func (p *PlayerController) PlayersInArea(x1, y1, z1 float64) (inArea []*PlayerController, notInArea []*PlayerController) {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.Players {
		if pl.UUID == p.UUID {
			continue
		}
		x2, y2, z2 := pl.player.Position()
		distance := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1))
		if float64(pl.ClientSettings().ViewDistance)*16 < distance {
			notInArea = append(notInArea, pl)
		} else {
			inArea = append(inArea, pl)
		}
	}
	return inArea, notInArea
}

func angle(x0, y0, z0 float64) (yaw float64, pitch float64) {
	dx := x0 - x0
	dy := y0 - y0
	dz := z0 - z0
	r := math.Sqrt(dx*dx + dy*dy + dz*dz)
	yaw = -math.Atan2(dx, dz) / math.Pi * 180
	if yaw < 0 {
		yaw = 360 + yaw
	}
	pitch = -math.Asin(dy/r) / math.Pi * 180
	return
}

func (p *PlayerController) BroadcastMovement(id int32, x1, y1, z1 float64, yaw, pitch float32, ong bool) {
	p.CalculateUnusedChunks()
	oldx, oldy, oldz := p.player.Position()
	p.player.SetPosition(x1, y1, z1, yaw, pitch, ong)
	inArea, notInArea := p.PlayersInArea(x1, y1, z1)

	for _, pl := range notInArea {
		if pl.IsSpawned(p.player.EntityId()) {
			pl.DespawnPlayer(p)
		}
	}
	for _, pl := range inArea {
		if pl.IsSpawned(p.player.EntityId()) {
			switch id {
			case 0x14: // position
				pl.session.SendPacket(&packet.EntityPosition{
					EntityID: p.player.EntityId(),
					X:        int16(((x1 * 32) - oldx*32) * 128),
					Y:        int16(((y1 * 32) - oldy*32) * 128),
					Z:        int16(((z1 * 32) - oldz*32) * 128),
					OnGround: ong,
				})
			case 0x15: // position + rotation
				yaw, pitch := angle(x1, y1, z1)
				pl.session.SendPacket(&packet.EntityPositionRotation{
					EntityID: p.player.EntityId(),
					X:        int16(((x1 * 32) - oldx*32) * 128),
					Y:        int16(((y1 * 32) - oldy*32) * 128),
					Z:        int16(((z1 * 32) - oldz*32) * 128),
					Yaw:      byte(yaw),
					Pitch:    byte(pitch),
					OnGround: ong,
				})
				pl.session.SendPacket(&packet.EntityHeadRotation{
					EntityID: p.player.EntityId(),
					HeadYaw:  uint8(yaw),
				})
			case 0x16: // rotation
				//yaw, pitch := angle(x1, y1, z1)
				pl.session.SendPacket(&packet.EntityRotation{
					EntityID: p.player.EntityId(),
					Yaw:      byte(yaw),
					Pitch:    byte(pitch),
					OnGround: ong,
				})
				pl.session.SendPacket(&packet.EntityHeadRotation{
					EntityID: p.player.EntityId(),
					HeadYaw:  uint8(yaw),
				})
			}
		} else {
			pl.SpawnPlayer(p)
		}
	}
}

func (p *PlayerController) BroadcastPose(pose int32) {
	inArea, _ := p.PlayersInArea(p.Position())
	for _, pl := range inArea {
		pl.session.SendPacket(&PacketSetPose{EntityID: p.player.EntityId(), Pose: pose})
	}
}

func (p *PlayerController) BroadcastSprinting(val bool) {
	//inArea, _ := p.PlayersInArea(p.Position())
	//for _, pl := range inArea {
	//	//pl.session.SendPacket(&PacketSetPose{EntityID: p.player.EntityId(), Pose: pose})
	//}
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

type PacketSetPose struct {
	EntityID int32
	Pose     int32
}

func (*PacketSetPose) ID() int32 {
	return 0x52
}

func (*PacketSetPose) Decode(*packet.Reader) error {
	return nil
}

func (s PacketSetPose) Encode(w packet.Writer) error {
	w.VarInt(s.EntityID)
	w.Uint8(6)
	w.VarInt(20)
	w.VarInt(s.Pose)
	return w.Uint8(0xFF)
}
