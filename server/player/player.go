package player

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/network"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/google/uuid"
)

type Player struct {
	UUID    string
	Session *network.Session
}

func NewPlayer(s *network.Session) *Player {
	uuid, _ := uuid.FromBytes(s.Conn.Info.UUID[:])
	return &Player{Session: s, UUID: uuid.String()}
}

func (p *Player) SetCommands(commands *packet.DeclareCommands) error {
	return p.Session.Conn.SendPacket(commands)
}

func (p *Player) SystemChatMessage(content string) error {
	return p.Session.Conn.SendPacket(&packet.SystemChatMessage{Content: content})
}

func (p *Player) JoinDimension(eid int32, hardcore bool, gm byte, d *world.Dimension, seed int64, vd, sd int32) error {
	if err := p.Session.Conn.SendPacket(&packet.JoinGame{
		EntityID:           eid,
		IsHardcore:         hardcore,
		GameMode:           gm,
		PreviousGameMode:   -1,
		DimensionNames:     []string{d.Type()},
		DimensionName:      d.Type(),
		DimensionType:      d.Type(),
		HashedSeed:         seed,
		ViewDistance:       vd,
		SimulationDistance: sd,
		PartialCooldown:    3,
	}); err != nil {
		return err
	}

	if err := p.Session.Conn.SendPacket(packet.PluginMessage{
		Channel: "minecraft:brand",
		Data:    []byte("DynamiteMC"),
	}); err != nil {
		return err
	}

	if err := p.Session.Conn.SendPacket(&packet.SetDefaultSpawnPosition{}); err != nil {
		return err
	}
	return nil
}
