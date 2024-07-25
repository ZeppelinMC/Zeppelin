package terrain

import (
	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/zeppelinmc/zeppelin/server/world/region"
)

type TerrainGenerator struct {
	noise *perlin.Perlin
}

func (g TerrainGenerator) NewChunk(cx, cz int32) region.Chunk {
	c := region.NewChunk(cx, cz)

	for x := int32(0); x < 16; x++ {
		for z := int32(0); z < 16; z++ {
			absX, absZ := cx*16+x, cz*16+z

			y := int32(g.noise.Noise2D(float64(absX)/25, float64(absZ)/25) * 50)

			if y <= -64 {
				y = -64
			} else if y >= 320 {
				y = 320
			}

			c.SetBlock(x, y, z, grassBlock)
			for s := int32(region.MinChunkY * 16); s < y-1; s++ {
				c.SetBlock(x, int32(s), z, grassBlock)
			}
		}
	}

	return c
}

var grassBlock = region.Block{Name: "minecraft:grass_block", Properties: map[string]string{"snowy": "false"}}

func init() {
	region.Def = TerrainGenerator{
		noise: perlin.NewPerlin(2, 2, 1, rand.Int63()),
	}
}
