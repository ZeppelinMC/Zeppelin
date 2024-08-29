package entity

import (
	"sync/atomic"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/metadata"
)

var entityId atomic.Int32

func NewEntityId() int32 {
	return entityId.Add(1)
}

// The interface all entities should implement
type Entity interface {
	// The type of entity this is
	Type() int32
	// The unique global identifier of this entity
	UUID() uuid.UUID
	// The unique identifier of this entity for this server
	EntityId() int32
	// The position of this entity (x, y, z)
	Position() (x, y, z float64)
	// The rotation of this entity (yaw, pitch)
	Rotation() (yaw, pitch float32)
	// The metadata of this entity
	Metadata() metadata.Metadata
	// The attributes of this entity
	Attributes() []Attribute
	// The name of the dimension this entity is in
	Dimension() string
}

// The interface all living entities should implement
type LivingEntity interface {
	Entity
	// the health of this entity
	Health() float32
	// the food level of this entity
	Food() int32
	// the food saturation level of this entity
	FoodSaturation() float32
	// the food exhaustion level of this entity
	FoodExhaustion() float32
}

type Attribute struct {
	Base float64 `nbt:"base"`
	Id   string  `nbt:"id"`
}

var DefaultAttributes = map[string]float64{
	"minecraft:generic.armor":                     0,
	"minecraft:generic.armor_toughness":           0,
	"minecraft:generic.attack_damage":             2,
	"minecraft:generic.attack_knockback":          0,
	"minecraft:generic.attack_speed":              2,
	"minecraft:generic.block_break_speed":         1,
	"minecraft:generic.block_interaction_range":   4.5,
	"minecraft:generic.entity_interaction_range":  3,
	"minecraft:generic.fall_damage_multiplier":    1,
	"minecraft:generic.flying_speed":              0.4,
	"minecraft:generic.follow_range":              32,
	"minecraft:generic.gravity":                   0.08,
	"minecraft:generic.jump_strength":             0.42,
	"minecraft:generic.knockback_resistance":      0,
	"minecraft:generic.luck":                      0,
	"minecraft:generic.max_absorption":            0,
	"minecraft:generic.max_health":                20,
	"minecraft:generic.movement_speed":            0.7,
	"minecraft:generic.safe_fall_distance":        3,
	"minecraft:generic.scale":                     1,
	"minecraft:zombie.spawn_reinforcements":       0,
	"minecraft:generic.step_height":               0.6,
	"minecraft:generic.submerged_mining_speed":    0.2,
	"minecraft:generic.sweeping_damage_ratio":     0,
	"minecraft:generic.water_movement_efficiency": 0,
}
