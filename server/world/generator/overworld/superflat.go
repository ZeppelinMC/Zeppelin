package overworld

import (
	"github.com/dynamitemc/dynamite/server/block"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type FlatGenerator struct{}

func (f *FlatGenerator) GenerateChunk(x, z int32) (*chunk.Chunk, error) {
	c := chunk.Chunk{}
	for x := int64(0); x < 16; x++ {
		for z := int64(0); z < 16; z++ {
			c.SetBlock(x, -64, z, block.Bedrock{})
			c.SetBlock(x, -63, z, block.Dirt{})
			c.SetBlock(x, -62, z, block.Dirt{})
			c.SetBlock(x, -61, z, block.GrassBlock{})
		}
	}
	return nil, nil
}
