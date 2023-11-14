package generator

import (
	"github.com/dynamitemc/dynamite/server/block"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type FlatGenerator struct{}

func (f *FlatGenerator) GenerateChunk(x, z int32) (*chunk.Chunk, error) {
	c := chunk.Chunk{}
	for x := int64(0); x < 16; x++ {
		for z := int64(0); z < 16; z++ {
			c.SetBlock(x, 0, z, block.Bedrock{})
			c.SetBlock(x, 1, z, block.Dirt{})
			c.SetBlock(x, 2, z, block.Dirt{})
			c.SetBlock(x, 3, z, block.GrassBlock{})
		}
	}
	return nil, nil
}
