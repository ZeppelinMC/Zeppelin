package entity

import (
	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/net/metadata"
)

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
