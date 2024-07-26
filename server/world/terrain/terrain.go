package terrain

import (
	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/zeppelinmc/zeppelin/server/world/region"
)

type TerrainGenerator struct {
	noise *perlin.Perlin
	seed  int64
}

func (g TerrainGenerator) NewChunk(cx, cz int32) region.Chunk {
	c := region.NewChunk(cx, cz)

	var (
		treeY int32
	)

	for x := int32(0); x < 16; x++ {
		for z := int32(0); z < 16; z++ {
			absX, absZ := cx*16+x, cz*16+z

			y := int32(g.noise.Noise2D(float64(absX)/30, float64(absZ)/30)*10) + 80

			if y <= -64 {
				y = -64
			} else if y >= 320 {
				y = 320
			}

			c.SetBlock(x, y, z, grassBlock)
			for s := int32(region.MinChunkY * 16); s < y; s++ {
				c.SetBlock(x, int32(s), z, dirt)
			}

			if x == 7 && z == 7 {
				treeY = y
			}
		}
	}

	if rand.Int31()&0x01 == 00 {
		g.generateTree(c, 7, 7, treeY)
	}

	return c
}

func (g TerrainGenerator) generateTree(c region.Chunk, x, z, surface int32) {
	c.SetBlock(x, surface+1, z, oakLog)
	c.SetBlock(x, surface+2, z, oakLog)
	c.SetBlock(x, surface+3, z, oakLog)
	c.SetBlock(x, surface+4, z, oakLog)

	for i := x - 2; i <= x+2; i++ {
		for j := z - 2; j <= z+2; j++ {
			if isborder(i, j, x-2, z-2, x+2, z+2) {
				continue
			}
			c.SetBlock(i, surface+5, j, oakLeaves)
			if i != x || j != z {
				c.SetBlock(i, surface+4, j, oakLeaves)
			}
		}
	}
	for i := x - 1; i <= x+1; i++ {
		for j := z - 1; j <= z+1; j++ {
			c.SetBlock(i, surface+6, j, oakLeaves)

			if !isborder(i, j, x-1, z-1, x+1, z+1) {
				c.SetBlock(i, surface+7, j, oakLeaves)
			}
		}
	}
}

func NewTerrainGenerator(seed int64) TerrainGenerator {
	return TerrainGenerator{
		noise: perlin.NewPerlin(2, 2, 1, seed),
		seed:  seed,
	}
}

func isborder(x, z, minX, minZ, maxX, maxZ int32) bool {
	return (x == minX && z == minZ) || (x == maxX && z == maxZ) || (x == maxX && z == minZ) || (x == minZ && z == maxZ)
}

var grassBlock = region.Block{Name: "minecraft:grass_block", Properties: map[string]string{"snowy": "false"}}
var dirt = region.Block{Name: "minecraft:dirt"}
var bedrock = region.Block{Name: "minecraft:bedrock"}
var oakLog = region.Block{Name: "minecraft:oak_log", Properties: map[string]string{"axis": "y"}}
var oakLeaves = region.Block{Name: "minecraft:oak_leaves", Properties: map[string]string{
	"distance":    "1",
	"persistent":  "false",
	"waterlogged": "false",
}}
