package server

import (
	"math"
	"strings"
	"time"

	"github.com/aimjel/minecraft/protocol/types"

	"github.com/aimjel/minecraft/packet"
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

func (p *PlayerController) BreakBlock(pos uint64) {
	p.Server.GlobalBroadcast(&packet.BlockUpdate{
		Location: int64(pos),
	})
}

func (p *PlayerController) BroadcastDigging(pos uint64) {
	i := byte(0)
	id := p.player.EntityId()
	in, _ := p.PlayersInArea(p.Position())
	for range time.NewTicker(time.Millisecond * 100).C {
		if i > 10 {
			break
		}
		for _, pl := range in {
			pl.session.SendPacket(&packet.SetBlockDestroyStage{
				EntityID:     id,
				Location:     pos,
				DestroyStage: i,
			})
		}
		i++
	}
}

func (p *PlayerController) BroadcastSkinData() {
	cl := p.ClientSettings()
	p.Server.GlobalBroadcast(&PacketSetPlayerMetadata{
		EntityID:           p.player.EntityId(),
		DisplayedSkinParts: &cl.DisplayedSkinParts,
		MainHand:           &cl.MainHand,
	})
}

func degreesToAngle(degrees float32) byte {
	return byte(math.Round(float64(degrees) * (256.0 / 360.0)))
}

func positionIsValid(x, y, z float64) bool {
	return !math.IsNaN(x) && !math.IsNaN(y) && !math.IsNaN(z) &&
		!math.IsInf(x, 0) && !math.IsInf(y, 0) && !math.IsInf(z, 0)
}

func direction(ya, pi float32) (x, y, z float64) {
	yaw, pitch := float64(ya), float64(pi)
	x = -math.Cos(pitch) * math.Sin(yaw)
	y = -math.Sin(pitch)
	z = math.Cos(pitch) * math.Cos(yaw)
	return x, y, z
}

func (p *PlayerController) Hit(entityId int32) {
	e := p.Server.FindEntity(entityId)
	x, y, z := p.Position()
	//yaw, pitch := p.Rotation()
	//d := direction(yaw, pitch)
	if pl, ok := e.(*PlayerController); ok {
		if pl.GameMode() == 1 {
			return
		}
		health := pl.player.Health()
		pl.SetHealth(health - 1)
		x1, y1, z1 := pl.Position()
		/*switch d {
		case 0:
			x1 += 0.5
			z1 += 0.5
		case 1:
			x1 += 0.5
			z1 -= 0.5
		case 2:
			x1 -= 0.5
			z1 += 0.5
		case 3:
			x1 -= 0.5
			z1 -= 0.5
		}*/
		pl.Push(x1, y1, z1)
	}
	p.BroadcastPacketAll(&packet.DamageEvent{
		EntityID:        entityId,
		SourceTypeID:    1,
		SourceCauseID:   p.player.EntityId() + 1,
		SourceDirectID:  p.player.EntityId() + 1,
		SourcePositionX: &x,
		SourcePositionY: &y,
		SourcePositionZ: &z,
	})
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
	oldx, oldy, oldz := p.player.Position()
	distance := math.Sqrt((x1-oldx)*(x1-oldx) + (y1-oldy)*(y1-oldy) + (z1-oldz)*(z1-oldz))
	if distance > 100 && !teleport {
		p.Teleport(oldx, oldy, oldz, yaw, pitch)
		p.Server.Logger.Info("%s moved too quickly!", p.Name())
		return
	}
	if !positionIsValid(x1, y1, z1) {
		p.Disconnect("Invalid move player packet received")
		return
	}

	p.player.SetPosition(x1, y1, z1, yaw, pitch, ong)
	inArea, notInArea := p.PlayersInArea(x1, y1, z1)

	if distance > 8 {
		id = 0
	}

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
					HeadYaw:  yaw,
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
					HeadYaw:  yaw,
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
		pl.session.SendPacket(&PacketSetPlayerMetadata{EntityID: p.player.EntityId(), Pose: &pose})
	}
}

func (p *PlayerController) BroadcastPacketAll(pk packet.Packet) {
	inArea, _ := p.AllPlayersInArea(p.Position())
	for _, pl := range inArea {
		pl.session.SendPacket(pk)
	}
}

func (p *PlayerController) BroadcastHealth() {
	inArea, _ := p.PlayersInArea(p.Position())
	h := p.player.Health()
	for _, pl := range inArea {
		pl.session.SendPacket(&PacketSetPlayerMetadata{EntityID: p.player.EntityId(), Health: &h})
	}
}

func (p *PlayerController) BroadcastSprinting(val bool) {
	inArea, _ := p.PlayersInArea(p.Position())
	for _, pl := range inArea {
		data := byte(0)
		if val {
			data |= 0x08
		}
		pl.session.SendPacket(&PacketSetPlayerMetadata{EntityID: p.player.EntityId(), Data: &data})
	}
}

func (srv *Server) PlayerlistUpdate() {
	players := make([]types.PlayerInfo, 0, len(srv.Players))
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		players = append(players, types.PlayerInfo{
			UUID:       p.session.conn.UUID(),
			Name:       p.session.conn.Name(),
			Properties: p.session.conn.Properties(),
			Listed:     true,
		})
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

type PacketSetPlayerMetadata struct {
	EntityID           int32
	Pose               *int32
	Data               *byte
	Health             *float32
	DisplayedSkinParts *uint8
	MainHand           *int32
}

func (*PacketSetPlayerMetadata) ID() int32 {
	return 0x52
}

func (*PacketSetPlayerMetadata) Decode(*packet.Reader) error {
	return nil
}

func (s PacketSetPlayerMetadata) Encode(w packet.Writer) error {
	w.VarInt(s.EntityID)
	if s.Pose != nil {
		w.Uint8(6)
		w.VarInt(20)
		w.VarInt(*s.Pose)
	}
	if s.Data != nil {
		w.Uint8(0)
		w.Uint8(0)
		w.Uint8(*s.Data)
	}
	if s.Health != nil {
		w.Uint8(9)
		w.VarInt(1)
		w.Float32(*s.Health)
	}
	if s.DisplayedSkinParts != nil {
		w.Uint8(17)
		w.VarInt(0)
		w.Uint8(*s.DisplayedSkinParts)
	}
	if s.MainHand != nil {
		w.Uint8(18)
		w.VarInt(0)
		w.Uint8(uint8(*s.MainHand))
	}
	return w.Uint8(0xFF)
}
