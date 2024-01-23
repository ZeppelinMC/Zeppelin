package generator

import (
	"github.com/aimjel/nitrate/server/block"
	"github.com/aimjel/nitrate/server/world/chunk"
)

type Flat struct{}

func (f Flat) GenerateChunk(x, z int32) *chunk.Chunk {
	c := chunk.New(x, z)

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			c.SetBlock(i, -64, j, block.Bedrock{})
			c.SetBlock(i, -63, j, block.Dirt{})
			c.SetBlock(i, -62, j, block.Dirt{})
			c.SetBlock(i, -61, j, block.Dirt{})
			c.SetBlock(i, -60, j, block.GrassBlock{})

			c.HeightMap.SetWorldSurface(i, j, -60)
			c.HeightMap.SetMotionBlocking(i, j, -60)
		}
	}

	c.GenerateSkyLight()

	return c
}
