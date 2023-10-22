package server

import (
	"fmt"
	"math"
	"strings"

	"github.com/aimjel/minecraft/protocol/types"
	"github.com/dynamitemc/dynamite/server/registry"

	"github.com/aimjel/minecraft/packet"
)

func (srv *Server) GlobalBroadcast(pk packet.Packet) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		p.SendPacket(pk)
	}
}

func (srv *Server) GlobalMessage(message string, sender *Session) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	for _, p := range srv.Players {
		if p.clientInfo.ChatMode == 2 {
			continue
		} else if p.clientInfo.ChatMode == 1 && sender != nil {
			continue
		}
		p.SendPacket(&packet.SystemChatMessage{
			Content: message,
		})
	}
	srv.Logger.Print(message)
}

func (srv *Server) OperatorMessage(message string) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if p.clientInfo.ChatMode == 2 || !p.player.Operator() {
			continue
		}
		p.SendPacket(&packet.SystemChatMessage{
			Content: message,
		})
	}
	message = strings.ReplaceAll(message, "§", "&")
	srv.Logger.Print(message)
}

func (p *Session) BroadcastAnimation(animation uint8) {
	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()

	for _, pl := range p.Server.Players {
		if !pl.IsSpawned(p.entityID) {
			continue
		}

		pl.SendPacket(&packet.EntityAnimation{
			EntityID:  p.entityID,
			Animation: animation,
		})
	}
}

func (p *Session) BreakBlock(pos uint64) {
	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()
	for _, pl := range p.Server.Players {
		if !pl.IsSpawned(p.entityID) {
			continue
		}
		pl.SendPacket(&packet.WorldEvent{Event: 2001, Location: pos})
		pl.SendPacket(&packet.BlockUpdate{
			Location: int64(pos),
		})
	}
}

/*func (p *Session) BroadcastDigging(pos uint64) {
	i := byte(0)
	id := p.entityID
	in, _ := p.PlayersInArea(p.Position())
	for range time.NewTicker(time.Millisecond * 100).C {
		if i > 9 {
			break
		}
		for _, pl := range in {
			if !pl.playReady {
				continue
			}

			pl.SendPacket(&packet.SetBlockDestroyStage{
				EntityID:     id,
				Location:     pos,
				DestroyStage: i,
			})
		}
		i++
	}
}*/

func (p *Session) BroadcastSkinData() {
	cl := p.clientInfo
	p.Server.GlobalBroadcast(&PacketSetPlayerMetadata{
		EntityID:           p.entityID,
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

func (p *Session) Hit(entityId int32) {
	e := p.Server.FindEntity(entityId)
	x, y, z := p.Position()
	soundId := int32(519)
	if pl, ok := e.(*Session); ok {
		if pl.GameMode() == 1 {
			return
		}
		health := pl.player.Health()
		pl.SetHealth(health - 1)
		pl.SendPacket(&packet.DamageEvent{
			EntityID:        entityId,
			SourceTypeID:    1,
			SourceCauseID:   p.entityID + 1,
			SourceDirectID:  p.entityID + 1,
			SourcePositionX: &x,
			SourcePositionY: &y,
			SourcePositionZ: &z,
		})
	} else {
		entity := e.(*Entity)
		sound, ok := registry.GetSound(fmt.Sprintf("minecraft:entity.%s.hurt", strings.TrimPrefix(entity.data.Id, "minecraft:")))
		if ok {
			soundId = sound.ProtocolID
		}
	}

	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()
	for _, pl := range p.Server.Players {
		if !pl.IsSpawned(entityId) {
			continue
		}
		pl.SendPacket(&packet.DamageEvent{
			EntityID:        entityId,
			SourceTypeID:    1,
			SourceCauseID:   p.entityID + 1,
			SourceDirectID:  p.entityID + 1,
			SourcePositionX: &x,
			SourcePositionY: &y,
			SourcePositionZ: &z,
		})
		pl.SendPacket(&packet.EntitySoundEffect{
			Category: 8,
			SoundID:  soundId,
			EntityID: entityId,
			Seed:     p.Server.World.Seed(),
			Volume:   1,
			Pitch:    1,
		})

	}
}

func (p *Session) Despawn() {
	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()
	for _, pl := range p.Server.Players {
		if !pl.IsSpawned(p.entityID) {
			continue
		}
		pl.DespawnPlayer(p)
	}
}

// InView Checks if p can see pl
func (p *Session) InView(pl *Session) bool {
	if !pl.playReady || p.player.Dimension() != pl.player.Dimension() {
		return false
	}

	x1, y1, z1 := p.Position()
	x2, y2, z2 := pl.Position()
	distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))

	return float64(p.clientInfo.ViewDistance)*16 > distance
}

func (p *Session) BroadcastMovement(id int32, x1, y1, z1 float64, ya, pi float32, ong bool, teleport bool) {
	oldx, oldy, oldz := p.player.Position()
	p.player.SetPosition(x1, y1, z1, ya, pi, ong)
	distance := math.Sqrt((x1-oldx)*(x1-oldx) + (y1-oldy)*(y1-oldy) + (z1-oldz)*(z1-oldz))
	if distance > 100 && !teleport {
		//p.Teleport(oldx, oldy, oldz, yaw, pitch)
		p.Server.Logger.Info("%s moved too quickly!", p.Name())
		return
	}
	if !positionIsValid(x1, y1, z1) {
		p.Disconnect("Invalid move player packet received")
		return
	}

	if distance > 8 {
		id = 0
	}

	yaw, pitch := degreesToAngle(ya), degreesToAngle(pi)

	var pk packet.Packet
	headRotationPacket := &packet.EntityHeadRotation{
		EntityID: p.entityID,
		HeadYaw:  yaw,
	}
	var sendHeadRotation bool
	switch id {
	case 0x14: // position
		pk = &packet.EntityPosition{
			EntityID: p.entityID,
			X:        int16(((x1 * 32) - oldx*32) * 128),
			Y:        int16(((y1 * 32) - oldy*32) * 128),
			Z:        int16(((z1 * 32) - oldz*32) * 128),
			OnGround: ong,
		}
	case 0x15: // position + rotation
		pk = &packet.EntityPositionRotation{
			EntityID: p.entityID,
			X:        int16(((x1 * 32) - oldx*32) * 128),
			Y:        int16(((y1 * 32) - oldy*32) * 128),
			Z:        int16(((z1 * 32) - oldz*32) * 128),
			Yaw:      yaw,
			Pitch:    pitch,
			OnGround: ong,
		}
		sendHeadRotation = true
	case 0x16: // rotation
		pk = &packet.EntityRotation{
			EntityID: p.entityID,
			Yaw:      yaw,
			Pitch:    pitch,
			OnGround: ong,
		}
		sendHeadRotation = true
	default:
		pk = &packet.TeleportEntity{
			EntityID: p.entityID,
			X:        x1,
			Y:        y1,
			Z:        z1,
			Yaw:      yaw,
			Pitch:    pitch,
			OnGround: ong,
		}
	}

	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()

	for _, pl := range p.Server.Players {
		if p.UUID == pl.UUID {
			continue
		}
		if !pl.InView(p) {
			pl.DespawnPlayer(p)
			continue
		}

		if !pl.IsSpawned(p.entityID) {
			pl.SpawnPlayer(p)
		}

		pl.SendPacket(pk)
		if sendHeadRotation {
			pl.SendPacket(headRotationPacket)
		}
	}
}

func (p *Session) BroadcastPose(pose int32) {
	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()
	pk := &PacketSetPlayerMetadata{
		EntityID: p.entityID,
		Pose:     &pose,
	}
	for _, pl := range p.Server.Players {
		if pl.IsSpawned(p.entityID) {
			pl.SendPacket(pk)
		}
	}
}

func (p *Session) BroadcastHealth() {
	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()
	h := p.player.Health()
	for _, pl := range p.Server.Players {
		if pl.IsSpawned(p.entityID) {
			pl.SendPacket(&PacketSetPlayerMetadata{EntityID: p.entityID, Health: &h})
		}
	}
}

func (p *Session) BroadcastSprinting(val bool) {
	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()

	data := byte(0)
	if val {
		data |= 0x08
	}

	pk := &PacketSetPlayerMetadata{EntityID: p.entityID, Data: &data}

	for _, pl := range p.Server.Players {
		if !pl.IsSpawned(p.entityID) {
			continue
		}
		pl.SendPacket(pk)
	}
}

func (srv *Server) PlayerlistUpdate() {
	players := make([]types.PlayerInfo, 0, len(srv.Players))
	srv.mu.Lock()
	for _, p := range srv.Players {
		players = append(players, types.PlayerInfo{
			UUID:       p.conn.UUID(),
			Name:       p.conn.Name(),
			Properties: p.conn.Properties(),
			Listed:     true,
		})
	}
	srv.mu.Unlock()
	srv.GlobalBroadcast(&packet.PlayerInfoUpdate{
		Actions: 0x01 | 0x08,
		Players: players,
	})
}

func (srv *Server) PlayerlistRemove(players ...[16]byte) {
	srv.GlobalBroadcast(&packet.PlayerInfoRemove{UUIDs: players})
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
