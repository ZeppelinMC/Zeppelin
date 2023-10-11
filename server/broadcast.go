package server

import (
	"math"
	"strings"

	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/player"
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
	srv.Logger.Print(message)
}

func (srv *Server) OperatorMessage(message string) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if p.ClientSettings().ChatMode == 2 || !p.player.Operator() {
			continue
		}
		p.session.SendPacket(&packet.SystemChatMessage{
			Content: message,
		})
	}
	message = strings.ReplaceAll(message, "§", "&")
	srv.Logger.Print(message)
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

func (p *PlayerController) AllPlayersInArea(x1, y1, z1 float64) (inArea []*PlayerController, notInArea []*PlayerController) {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.Players {
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

func (p *PlayerController) BroadcastAnimation(animation uint8) {
	inarea, _ := p.PlayersInArea(p.Position())
	id := p.player.EntityId()
	for _, pl := range inarea {
		pl.session.SendPacket(&packet.EntityAnimation{
			EntityID:  id,
			Animation: animation,
		})
	}
}

func degreesToAngle(degrees float32) byte {
	return byte(math.Round(float64(degrees) * (256.0 / 360.0)))
}

func positionIsValid(x, y, z float64) bool {
	return !math.IsNaN(x) && !math.IsNaN(y) && !math.IsNaN(z) &&
		!math.IsInf(x, 0) && !math.IsInf(y, 0) && !math.IsInf(z, 0)
}

func (p *PlayerController) Hit(entityId int32) {
	pl := p.Server.FindPlayerByID(entityId)
	if pl.GameMode() == 1 {
		return
	}
	//p.BroadcastPose(4)
	health := pl.player.Health()
	pl.SetHealth(health - 1)
	inarea, _ := p.AllPlayersInArea(p.Position())
	yaw, _ := p.Rotation()
	for _, v := range inarea {
		v.session.SendPacket(&packet.HurtAnimation{
			EntityID: pl.player.EntityId(),
			Yaw:      yaw,
		})
	}
}

func (p *PlayerController) Despawn() {
	inArea, _ := p.PlayersInArea(p.Position())
	for _, pl := range inArea {
		if pl.IsSpawned(p.player.EntityId()) {
			pl.DespawnPlayer(p)
		}
	}
}

func (p *PlayerController) BroadcastMovement(id int32, x1, y1, z1 float64, yaw, pitch float32, ong bool, teleport bool) {
	if !teleport {
		p.CalculateUnusedChunks()
	}
	oldx, oldy, oldz := p.player.Position()
	distance := math.Sqrt((x1-oldx)*(x1-oldx) + (y1-oldy)*(y1-oldy) + (z1-oldz)*(z1-oldz))
	if distance > 100 && !teleport {
		p.Disconnect("You moved too quickly :( (Hacking?)")
		return
	}
	if !positionIsValid(x1, y1, z1) {
		p.Disconnect("Illegal position")
		return
	}

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
				yaw, pitch := degreesToAngle(yaw), degreesToAngle(pitch)
				pl.session.SendPacket(&packet.EntityPositionRotation{
					EntityID: p.player.EntityId(),
					X:        int16(((x1 * 32) - oldx*32) * 128),
					Y:        int16(((y1 * 32) - oldy*32) * 128),
					Z:        int16(((z1 * 32) - oldz*32) * 128),
					Yaw:      yaw,
					Pitch:    pitch,
					OnGround: ong,
				})
				pl.session.SendPacket(&packet.EntityHeadRotation{
					EntityID: p.player.EntityId(),
					HeadYaw:  uint8(yaw),
				})
			case 0x16: // rotation
				yaw, pitch := degreesToAngle(yaw), degreesToAngle(pitch)
				pl.session.SendPacket(&packet.EntityRotation{
					EntityID: p.player.EntityId(),
					Yaw:      yaw,
					Pitch:    pitch,
					OnGround: ong,
				})
				pl.session.SendPacket(&packet.EntityHeadRotation{
					EntityID: p.player.EntityId(),
					HeadYaw:  uint8(yaw),
				})
			default:
				yaw, pitch := degreesToAngle(yaw), degreesToAngle(pitch)

				pl.session.SendPacket(&packet.TeleportEntity{
					EntityID: p.player.EntityId(),
					X:        x1,
					Y:        y1,
					Z:        z1,
					Yaw:      yaw,
					Pitch:    pitch,
					OnGround: ong,
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

func (p *PlayerController) BroadcastHealth() {
	inArea, _ := p.PlayersInArea(p.Position())
	for _, pl := range inArea {
		pl.session.SendPacket(&PacketSetHealth{EntityID: p.player.EntityId(), Health: p.player.Health()})
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

type PacketSetHealth struct {
	EntityID int32
	Health   float32
}

func (*PacketSetHealth) ID() int32 {
	return 0x52
}

func (*PacketSetHealth) Decode(*packet.Reader) error {
	return nil
}

func (s PacketSetHealth) Encode(w packet.Writer) error {
	w.VarInt(s.EntityID)
	w.Uint8(9)
	w.VarInt(1)
	w.Float32(s.Health)
	return w.Uint8(0xFF)
}
