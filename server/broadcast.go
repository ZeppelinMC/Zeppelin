package server

import (
	"fmt"
	"math"
	"strings"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/protocol/types"
	"github.com/dynamitemc/dynamite/server/item"
	"github.com/dynamitemc/dynamite/server/registry"
	"github.com/dynamitemc/dynamite/server/world"

	"github.com/aimjel/minecraft/packet"
)

func (srv *Server) GlobalBroadcast(pk packet.Packet) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.players {
		p.SendPacket(pk)
	}
}

func (srv *Server) GlobalMessage(message chat.Message) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.players {
		if p.clientInfo.ChatMode == 2 {
			continue
		}
		p.SendPacket(&packet.SystemChatMessage{
			Message: message,
		})
	}
	srv.Logger.Print(message)
}

func (srv *Server) OperatorMessage(message string) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	msg := chat.NewMessage(message)
	for _, p := range srv.players {
		if p.clientInfo.ChatMode == 2 || !p.Player.Operator() {
			continue
		}
		p.SendPacket(&packet.SystemChatMessage{
			Message: msg,
		})
	}
	srv.Logger.Print(msg)
}

func (p *Session) BroadcastAnimation(animation uint8) {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()

	for _, pl := range p.Server.players {
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
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	p.Server.GetDimension(p.Player.Dimension()).Block(world.ParsePosition(pos))
	for _, pl := range p.Server.players {
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
	x, y, z := p.Player.Position()
	soundId := int32(519)
	if pl, ok := e.(*Session); ok {
		if pl.Player.GameMode() == 1 {
			return
		}
		health := pl.Player.Health()
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

	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.players {
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
			Seed:     world.RandomSeed(),
			Volume:   1,
			Pitch:    1,
		})
	}
}

func (p *Session) Despawn() {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.players {
		if !pl.IsSpawned(p.entityID) {
			continue
		}
		pl.DespawnPlayer(p)
	}
}

// InView Checks if p can see pl
func (p *Session) InView(pl *Session) bool {
	if p.Player.Dimension() != pl.Player.Dimension() {
		return false
	}

	x1, y1, z1 := p.Player.Position()
	x2, y2, z2 := pl.Player.Position()
	distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))

	return float64(p.clientInfo.ViewDistance)*16 > distance
}

func (p *Session) BroadcastMovement(id int32, x1, y1, z1 float64, ya, pi float32, ong bool, teleport bool) {
	oldx, oldy, oldz := p.Player.Position()
	p.Player.SetPosition(x1, y1, z1, ya, pi, ong)
	distance := math.Sqrt((x1-oldx)*(x1-oldx) + (y1-oldy)*(y1-oldy) + (z1-oldz)*(z1-oldz))
	if distance > 100 && !teleport {
		//p.Teleport(oldx, oldy, oldz, yaw, pitch)
		p.Server.Logger.Info("%s moved too quickly!", p.Name())
		return
	}
	if !positionIsValid(x1, y1, z1) {
		p.Disconnect(p.Server.Translate("disconnect.invalid_player_movement", nil))
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

	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()

	for _, pl := range p.Server.players {
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

func (p *Session) BroadcastGamemode() {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	gm := int32(p.Player.GameMode())
	for _, sesh := range p.Server.players {
		sesh.SendPacket(&packet.PlayerInfoUpdate{
			Actions: 0x04,
			Players: []types.PlayerInfo{{
				UUID:     p.conn.UUID(),
				GameMode: gm,
			}},
		})
	}
}

func (p *Session) BroadcastPose(pose int32) {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	pk := &PacketSetPlayerMetadata{
		EntityID: p.entityID,
		Pose:     &pose,
	}
	for _, pl := range p.Server.players {
		if pl.IsSpawned(p.entityID) {
			pl.SendPacket(pk)
		}
	}
}

func (p *Session) BroadcastHealth() {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	h := p.Player.Health()
	for _, pl := range p.Server.players {
		if pl.IsSpawned(p.entityID) {
			pl.SendPacket(&PacketSetPlayerMetadata{EntityID: p.entityID, Health: &h})
		}
	}
}

func (p *Session) SendEquipment(pl *Session) {
	slots := make(map[int8]item.Item)
	inv := p.Player.Inventory()
	sel := p.Player.SelectedSlot()

	for _, s := range inv.Data() {
		switch s.Slot {
		case int8(sel):
			slots[0] = s
		case -106:
			slots[1] = s
		case 100:
			slots[2] = s
		case 101:
			slots[3] = s
		case 102:
			slots[4] = s
		case 103:
			slots[5] = s
		}
	}

	for s, i := range slots {
		pl.SendPacket(&SetEquipment{
			EntityID: p.entityID,
			Slot:     s,
			Item:     i,
		})
	}
}

func (p *Session) BroadcastEquipment() {
	slots := make(map[int8]item.Item)
	inv := p.Player.Inventory()
	sel := p.Player.SelectedSlot()

	for _, s := range inv.Data() {
		switch s.Slot {
		case int8(sel):
			slots[0] = s
		case -106:
			slots[1] = s
		case 100:
			slots[2] = s
		case 101:
			slots[3] = s
		case 102:
			slots[4] = s
		case 103:
			slots[5] = s
		}
	}

	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()

	for _, pl := range p.Server.players {
		if !pl.IsSpawned(p.entityID) {
			continue
		}
		for s, i := range slots {
			pl.SendPacket(&SetEquipment{
				EntityID: p.entityID,
				Slot:     s,
				Item:     i,
			})
		}
	}
}

func (p *Session) BroadcastSprinting(val bool) {
	data := byte(0)
	if val {
		data |= 0x08
	}

	pk := &PacketSetPlayerMetadata{EntityID: p.entityID, Data: &data}

	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()

	for _, pl := range p.Server.players {
		if !pl.IsSpawned(p.entityID) {
			continue
		}
		pl.SendPacket(pk)
	}
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
	Slot               *item.Item
	HandState          *int8
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
	if s.Slot != nil {
		sl, ok := registry.GetItem(s.Slot.Id)
		if ok {
			w.Uint8(8)
			w.Uint8(7)
			w.Bool(true)
			w.VarInt(sl.ProtocolID)
			w.Int8(s.Slot.Count)
			w.Nbt2(s.Slot.Tag)
		}
	}
	if s.HandState != nil {
		w.Uint8(8)
		w.Uint8(0)
		w.Int8(*s.HandState)
	}
	return w.Uint8(0xFF)
}

type SetEquipment struct {
	EntityID int32
	Slot     int8
	Item     item.Item
}

func (m SetEquipment) ID() int32 {
	return 0x55
}

func (m *SetEquipment) Decode(r *packet.Reader) error {
	return nil
}

func (m SetEquipment) Encode(w packet.Writer) error {
	w.VarInt(m.EntityID)
	w.Int8(m.Slot)
	id, ok := registry.GetItem(m.Item.Id)
	if !ok {
		w.Bool(false)
		return nil
	}
	w.Bool(true)
	w.VarInt(id.ProtocolID)
	w.Int8(m.Item.Count)
	w.Nbt2(m.Item.Tag)
	return nil
}
