package entity

import "github.com/zeppelinmc/zeppelin/server/world/level/uuid"

type LevelEntity struct {
	Air int16

	CustomName        string
	CustomNameVisible bool

	FallDistance  float32
	Fire          int16
	Glowing       bool
	HasVisualFire bool

	Id string `nbt:"id"`

	Invulnerable bool

	Motion     [3]float64
	NoGravity  bool
	OnGround   bool
	Passengers []Entity

	PortalCooldown int32
	Pos            [3]float64
	Rotation       [2]float32

	Silent      bool `nbt:"Silent,omitempty"`
	Tags        []string
	TicksFrozen int32 `nbt:"TicksFrozen,omitempty"`
	UUID        uuid.UUID

	//TODO add state subclasses
}
