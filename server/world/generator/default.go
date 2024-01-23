package generator

import (
	"github.com/aimjel/nitrate/server/block"
	"github.com/aimjel/nitrate/server/world/chunk"
	"github.com/aquilax/go-perlin"
)

type Default struct {
	Perlin *perlin.Perlin
}

func (d *Default) GenerateChunk(x, z int32) *chunk.Chunk {
	c := chunk.New(x, z)

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {

			worldX, worldZ := int(x)*16+i, int(z)*16+j

			y := int(d.Perlin.Noise2D(float64(worldX)/10, float64(worldZ)/10) * 5)

			if y <= -64 {
				y = -64
			} else if y >= 320 {
				y = 320
			}

			c.SetBlock(i, y, j, block.Stone{})

			for k := chunk.LowestY; k < y; k++ {
				c.SetBlock(i, k, j, block.Stone{})
			}

			c.HeightMap.SetWorldSurface(i, j, y)
			c.HeightMap.SetMotionBlocking(i, j, y)
		}
	}

	c.GenerateSkyLight()
	return c
}
