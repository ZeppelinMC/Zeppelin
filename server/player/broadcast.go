package player

import (
	"fmt"
	"math"
	"strings"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/protocol/types"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/controller"
	"github.com/dynamitemc/dynamite/server/entity"
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/registry"
	"github.com/dynamitemc/dynamite/server/world"

	"github.com/google/uuid"
)

func DegreesToAngle(degrees float32) byte {
	return byte(math.Round(float64(degrees) * (256.0 / 360.0)))
}

func PositionIsValid(x, y, z float64) bool {
	return !math.IsNaN(x) && !math.IsNaN(y) && !math.IsNaN(z) &&
		!math.IsInf(x, 0) && !math.IsInf(y, 0) && !math.IsInf(z, 0)
}

func (p *Player) BroadcastAnimation(animation uint8) {
	p.playerController.Range(func(u uuid.UUID, pl *Player) bool {
		if !pl.IsSpawned(p.EntityID()) || u == p.uuid {
			return true
		}

		pl.SendPacket(&packet.EntityAnimation{
			EntityID:  p.EntityID(),
			Animation: animation,
		})
		return true
	})
}

func (p *Player) BreakBlock(pos uint64) {
	//p.Player.Dimension().Block(world.ParsePosition(pos))
	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		if !pl.IsSpawned(p.EntityID()) {
			return true
		}

		pl.SendPacket(&packet.WorldEvent{Event: 2001, Location: pos})
		pl.SendPacket(&packet.BlockUpdate{
			Location: int64(pos),
		})
		return true
	})
}

func (p *Player) Attack(entityId int32) {
	e := findEntity(p.entityController, p.playerController, entityId)
	if e == nil {
		return
	}
	x, y, z := p.Position()
	soundId := int32(519)
	if pl, ok := e.(*Player); ok {
		if pl.GameMode() == enum.GameModeCreative {
			return
		}

		health := pl.Health()
		pl.SetHealth(health - 1)
		pl.SendPacket(&packet.DamageEvent{
			EntityID:        entityId,
			SourceTypeID:    enum.DamageTypePlayerAttack,
			SourceCauseID:   p.EntityID() + 1,
			SourceDirectID:  p.EntityID() + 1,
			SourcePositionX: &x,
			SourcePositionY: &y,
			SourcePositionZ: &z,
		})
	} else {
		entity, ok := e.(entity.LivingEntity)
		if !ok {
			return
		}
		sound, ok := registry.GetSound(fmt.Sprintf("minecraft:entity.%s.hurt", strings.TrimPrefix(entity.Type(), "minecraft:")))
		if ok {
			soundId = sound.ProtocolID
		}
	}

	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		if !pl.IsSpawned(entityId) {
			return true
		}
		pl.SendPacket(&packet.DamageEvent{
			EntityID:        entityId,
			SourceTypeID:    enum.DamageTypePlayerAttack,
			SourceCauseID:   p.EntityID() + 1,
			SourceDirectID:  p.EntityID() + 1,
			SourcePositionX: &x,
			SourcePositionY: &y,
			SourcePositionZ: &z,
		})
		pl.SendPacket(&packet.EntitySoundEffect{
			Category: enum.SoundCategoryAmbient,
			SoundID:  soundId,
			EntityID: entityId,
			Seed:     world.RandomSeed(),
			Volume:   1,
			Pitch:    1,
		})
		return true
	})
}

func (p *Player) Despawn() {
	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		if !pl.IsSpawned(p.EntityID()) {
			return true
		}
		pl.DespawnEntity(p.EntityID())
		return true
	})
}

func (p *Player) BroadcastMovement(id int32, x1, y1, z1 float64, ya, pi float32, ong bool, teleport bool) {
	oldx, oldy, oldz := p.Position()
	p.SetPosition(x1, y1, z1, ya, pi, ong)

	if g := p.GameMode(); g == enum.GameModeSurvival || g == enum.GameModeAdventure {
		if y1 > p.HighestY() {
			p.SetHighestY(y1)
		} else {
			d := p.HighestY() - y1
			if d > 3 {
				if b := p.OnBlock(); b != nil && b.EncodedName() != "minecraft:air" && b.EncodedName() != "minecraft:water" {
					d -= 3
					p.Damage(float32(d))
					p.SetHighestY(0)
				}
			}
		}
	}

	distance := math.Sqrt((x1-oldx)*(x1-oldx) + (y1-oldy)*(y1-oldy) + (z1-oldz)*(z1-oldz))
	/*if distance > 100 && !teleport {
		p.Teleport(oldx, oldy, oldz, yaw, pitch)
		p.logger.Info("%s moved too quickly!", p.Name())
		return
	}*/
	if !PositionIsValid(x1, y1, z1) {
		p.Disconnect(p.lang.Translate("disconnect.invalid_player_movement", nil))
		return
	}

	if distance > 8 {
		//id = 0
		return
	}

	yaw, pitch := DegreesToAngle(ya), DegreesToAngle(pi)

	var pk packet.Packet
	headRotationPacket := &packet.EntityHeadRotation{
		EntityID: p.EntityID(),
		HeadYaw:  yaw,
	}
	var sendHeadRotation bool
	switch id {
	case 0x14: // position
		pk = &packet.EntityPosition{
			EntityID: p.EntityID(),
			X:        int16(((x1 * 32) - oldx*32) * 128),
			Y:        int16(((y1 * 32) - oldy*32) * 128),
			Z:        int16(((z1 * 32) - oldz*32) * 128),
			OnGround: ong,
		}
	case 0x15: // position + rotation
		pk = &packet.EntityPositionRotation{
			EntityID: p.EntityID(),
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
			EntityID: p.EntityID(),
			Yaw:      yaw,
			Pitch:    pitch,
			OnGround: ong,
		}
		sendHeadRotation = true
	default:
		pk = &packet.TeleportEntity{
			EntityID: p.EntityID(),
			X:        x1,
			Y:        y1,
			Z:        z1,
			Yaw:      yaw,
			Pitch:    pitch,
			OnGround: ong,
		}
	}

	p.playerController.Range(func(u uuid.UUID, pl *Player) bool {
		if p.UUID() == u {
			return true
		}
		if !pl.InView(p.Position()) {
			pl.DespawnEntity(p.EntityID())
			return true
		}

		if !pl.IsSpawned(p.EntityID()) {
			pl.SpawnPlayer(p)
			return true
		}

		pl.SendPacket(pk)
		if sendHeadRotation {
			pl.SendPacket(headRotationPacket)
		}
		return true
	})
}

func (p *Player) BroadcastGamemode() {
	gm := int32(p.GameMode())

	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		pl.SendPacket(&packet.PlayerInfoUpdate{
			Actions: 0x04,
			Players: []types.PlayerInfo{{
				UUID:     p.conn.UUID(),
				GameMode: gm,
			}},
		})
		return true
	})
}

func (p *Player) sendEquipment(pl *Player) {
	slots := make(map[int8]packet.Slot)
	inv := p.Inventory
	sel := p.SelectedSlot()

	for _, s := range inv.Data() {
		switch s.Slot {
		case int8(sel):
			s, _ := s.ToPacketSlot()
			slots[0] = s
		case -106:
			s, _ := s.ToPacketSlot()
			slots[1] = s
		case 100:
			s, _ := s.ToPacketSlot()
			slots[2] = s
		case 101:
			s, _ := s.ToPacketSlot()
			slots[3] = s
		case 102:
			s, _ := s.ToPacketSlot()
			slots[4] = s
		case 103:
			s, _ := s.ToPacketSlot()
			slots[5] = s
		}
	}

	for s, i := range slots {
		pl.SendPacket(&packet.SetEquipment{
			EntityID: p.EntityID(),
			Slot:     s,
			Item:     i,
		})
	}
}

func (p *Player) BroadcastMetadataInArea(pk *packet.SetEntityMetadata) {
	p.playerController.Range(func(u uuid.UUID, pl *Player) bool {
		if !pl.IsSpawned(p.EntityID()) || u == p.uuid {
			return true
		}
		pl.SendPacket(pk)
		return true
	})
}

func (p *Player) broadcastMetadataGlobal(pk *packet.SetEntityMetadata) {
	p.playerController.Range(func(u uuid.UUID, pl *Player) bool {
		if u == p.uuid {
			return true
		}
		pl.SendPacket(pk)
		return true
	})
}

func findEntity(
	entities *controller.Controller[int32, entity.Entity],
	players *controller.Controller[uuid.UUID, *Player],
	id int32,
) interface{} {
	if _, p := players.Range(func(_ uuid.UUID, p *Player) bool {
		return p.entityID != id
	}); p != nil {
		return p
	} else {
		e, ok := entities.Get2(id)
		if !ok {
			return nil
		}
		return e
	}
}

func findEntityByUUID(
	entities *controller.Controller[int32, entity.Entity],
	players *controller.Controller[uuid.UUID, *Player],
	id [16]byte,
) interface{} {
	if _, p := players.Range(func(_ uuid.UUID, p *Player) bool {
		return p.UUID() != id
	}); p != nil {
		return p
	}
	if _, e := entities.Range(func(_ int32, e entity.Entity) bool {
		return e.UUID() != id
	}); e != nil {
		return e
	}
	return nil
}

func globalMessage(logger *logger.Logger, players *controller.Controller[uuid.UUID, *Player], message chat.Message) {
	players.Range(func(_ uuid.UUID, p *Player) bool {
		if p.ClientSettings().ChatMode == 2 {
			return true
		}
		p.SendPacket(&packet.SystemChatMessage{
			Message: message,
		})
		return true
	})
	logger.Print(message)
}
