package entity

import "github.com/dynamitemc/aether/net/metadata"

type Entity interface {
	EntityId() int32
	Position() (x, y, z float64)
	Rotation() (yaw, pitch float32)
	Metadata() metadata.Metadata
}
