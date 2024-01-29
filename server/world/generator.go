package world

import "github.com/dynamitemc/dynamite/server/world/chunk"

type Generator interface {
	GenerateChunk(x, z int32) *chunk.Chunk
}
