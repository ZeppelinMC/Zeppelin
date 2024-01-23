package world

import "github.com/aimjel/nitrate/server/world/chunk"

type Generator interface {
	GenerateChunk(x, z int32) *chunk.Chunk
}
