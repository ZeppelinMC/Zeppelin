package player

import (
	"encoding/binary"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/network"
	"github.com/dynamitemc/dynamite/server/world"
)

type Player struct {
	session *network.Session
}

func NewPlayer(s *network.Session) *Player {
	return &Player{session: s}
}

func (p *Player) JoinDimension(eid int32, hardcore bool, gm byte, d *world.Dimension, seed int64, vd, sd int32) error {
	hs := [8]byte{}
	if err := p.session.Conn.SendPacket(&packet.JoinGame{
		EntityID:           eid,
		IsHardcore:         hardcore,
		GameMode:           gm,
		PreviousGameMode:   -1,
		DimensionNames:     []string{d.Type()},
		DimensionName:      d.Type(),
		DimensionType:      d.Type(),
		HashedSeed:         int64(binary.BigEndian.Uint64(hs[:8])),
		ViewDistance:       vd,
		SimulationDistance: sd,
		PartialCooldown:    3,
	}); err != nil {
		return err
	}

	if err := p.session.Conn.SendPacket(&packet.SetDefaultSpawnPosition{}); err != nil {
		return err
	}
	return nil
}
