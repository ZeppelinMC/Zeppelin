package player

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/network"
	"github.com/dynamitemc/dynamite/server/world"
)

type BrandMessage string

func (s BrandMessage) ID() int32 {
	return 0x17
}

func (s BrandMessage) Decode(r *packet.Reader) error {
	return nil
}

func (s BrandMessage) Encode(w packet.Writer) error {
	w.String("minecraft:brand")
	return w.String(string(s))
}

type Player struct {
	Session *network.Session
}

func NewPlayer(s *network.Session) *Player {
	return &Player{Session: s}
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

	p.Session.Conn.SendPacket(BrandMessage("DynamiteMC"))

	if err := p.Session.Conn.SendPacket(&packet.SetDefaultSpawnPosition{}); err != nil {
		return err
	}
	return nil
}
