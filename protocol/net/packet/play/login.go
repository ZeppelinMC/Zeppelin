package play

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdLogin = 0x2B

type Login struct {
	EntityID int32

	Hardcore bool

	MaxPlayers int32

	ViewDistance, SimulationDistance int32

	ReducedDebugInfo,
	EnableRespawnScreen,
	DoLimitedCrafting bool

	Dimensions    []string
	DimensionType int32
	DimensionName string

	HashedSeed int64

	GameMode         byte
	PreviousGameMode int8

	IsDebug, IsFlat bool

	DeathDimensionName                             string
	DeathLocationX, DeathLocationY, DeathLocationZ int32

	PortalCooldown     int32
	EnforcesSecureChat bool
}

func (Login) ID() int32 {
	return PacketIdLogin
}

func (l *Login) Encode(w encoding.Writer) error {
	if err := w.Int(l.EntityID); err != nil {
		return err
	}
	if err := w.Bool(l.Hardcore); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(l.Dimensions))); err != nil {
		return err
	}
	for _, dim := range l.Dimensions {
		if err := w.Identifier(dim); err != nil {
			return err
		}
	}

	if err := w.VarInt(l.MaxPlayers); err != nil {
		return err
	}
	if err := w.VarInt(l.ViewDistance); err != nil {
		return err
	}
	if err := w.VarInt(l.SimulationDistance); err != nil {
		return err
	}
	if err := w.Bool(l.ReducedDebugInfo); err != nil {
		return err
	}
	if err := w.Bool(l.EnableRespawnScreen); err != nil {
		return err
	}
	if err := w.Bool(l.DoLimitedCrafting); err != nil {
		return err
	}
	if err := w.VarInt(l.DimensionType); err != nil {
		return err
	}
	if err := w.Identifier(l.DimensionName); err != nil {
		return err
	}
	if err := w.Long(l.HashedSeed); err != nil {
		return err
	}
	if err := w.Ubyte(l.GameMode); err != nil {
		return err
	}
	if err := w.Byte(l.PreviousGameMode); err != nil {
		return err
	}
	if err := w.Bool(l.IsDebug); err != nil {
		return err
	}
	if err := w.Bool(l.IsFlat); err != nil {
		return err
	}
	if err := w.Bool(l.DeathDimensionName != ""); err != nil {
		return err
	}
	if l.DeathDimensionName != "" {
		if err := w.Identifier(l.DeathDimensionName); err != nil {
			return err
		}
		if err := w.Position(l.DeathLocationX, l.DeathLocationY, l.DeathLocationZ); err != nil {
			return err
		}
	}
	if err := w.VarInt(l.PortalCooldown); err != nil {
		return err
	}
	return w.Bool(l.EnforcesSecureChat)
}

func (l *Login) Decode(r encoding.Reader) error {
	if err := r.Int(&l.EntityID); err != nil {
		return err
	}
	if err := r.Bool(&l.Hardcore); err != nil {
		return err
	}
	var dimlen int32
	if _, err := r.VarInt(&dimlen); err != nil {
		return err
	}
	if dimlen < 0 {
		return fmt.Errorf("negative length for make (login decode)")
	}
	l.Dimensions = make([]string, dimlen)
	for _, dim := range l.Dimensions {
		if err := r.Identifier(&dim); err != nil {
			return err
		}
	}

	if _, err := r.VarInt(&l.MaxPlayers); err != nil {
		return err
	}
	if _, err := r.VarInt(&l.ViewDistance); err != nil {
		return err
	}
	if _, err := r.VarInt(&l.SimulationDistance); err != nil {
		return err
	}
	if err := r.Bool(&l.ReducedDebugInfo); err != nil {
		return err
	}
	if err := r.Bool(&l.EnableRespawnScreen); err != nil {
		return err
	}
	if err := r.Bool(&l.DoLimitedCrafting); err != nil {
		return err
	}
	if _, err := r.VarInt(&l.DimensionType); err != nil {
		return err
	}
	if err := r.Identifier(&l.DimensionName); err != nil {
		return err
	}
	if err := r.Long(&l.HashedSeed); err != nil {
		return err
	}
	if err := r.Ubyte(&l.GameMode); err != nil {
		return err
	}
	if err := r.Byte(&l.PreviousGameMode); err != nil {
		return err
	}
	if err := r.Bool(&l.IsDebug); err != nil {
		return err
	}
	if err := r.Bool(&l.IsFlat); err != nil {
		return err
	}
	var hasDeathDim bool
	if err := r.Bool(&hasDeathDim); err != nil {
		return err
	}
	if hasDeathDim {
		if err := r.Identifier(&l.DeathDimensionName); err != nil {
			return err
		}
		if err := r.Position(&l.DeathLocationX, &l.DeathLocationY, &l.DeathLocationZ); err != nil {
			return err
		}
	}
	if _, err := r.VarInt(&l.PortalCooldown); err != nil {
		return err
	}
	return r.Bool(&l.EnforcesSecureChat)
}
