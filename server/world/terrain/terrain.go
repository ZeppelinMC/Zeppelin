package terrain

import (
	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/zeppelinmc/zeppelin/server/world/block"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type TerrainGenerator struct {
	noise *perlin.Perlin
	seed  int64
}

func (g TerrainGenerator) GenerateWorldSpawn() (x, y, z int32) {
	x, z = rand.Int31n(160), rand.Int31n(160)
	y = g.y(x, z)

	return
}

func (g TerrainGenerator) y(x, z int32) int32 {
	return int32(g.noise.Noise2D(float64(x)/30, float64(z)/30)*10) + 80
}

func (g TerrainGenerator) NewChunk(cx, cz int32) chunk.Chunk {
	c := chunk.NewChunk(cx, cz)

	var (
		treeY int32
	)

	for x := int32(0); x < 16; x++ {
		for z := int32(0); z < 16; z++ {
			absX, absZ := (cx<<4)+x, (cz<<4)+z

			y := g.y(absX, absZ)

			if y <= -64 {
				y = -64
			} else if y >= 320 {
				y = 320
			}

			c.Heightmaps.WorldSurface.Set(x, z, y)
			c.Heightmaps.MotionBlocking.Set(x, z, y)
			c.SetBlock(x, y, z, grassBlock)
			for s := int32(chunk.MinChunkY * 16); s < y; s++ {
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

// tree, woody plant that regularly renews its growth (perennial). Most plants classified as trees have a single self-supporting trunk containing woody tissues, and in most species the trunk produces secondary limbs, called branches
func (g TerrainGenerator) generateTree(c chunk.Chunk, x, z, surface int32) {
	// set the block under the tree to be dirt
	c.SetBlock(x, surface, z, dirt)
	// set 4 layers of oak logs
	c.SetBlock(x, surface+1, z, oakLog)
	c.SetBlock(x, surface+2, z, oakLog)
	c.SetBlock(x, surface+3, z, oakLog)
	c.SetBlock(x, surface+4, z, oakLog)

	for i := x - 2; i <= x+2; i++ {
		for j := z - 2; j <= z+2; j++ {
			// generate the bottom 2 layers of the tree's leaves (3x3 excluding corners)
			if !isCorner(i, j, x-2, z-2, x+2, z+2) {
				c.SetBlock(i, surface+5, j, oakLeaves)
				if i != x || j != z {
					c.SetBlock(i, surface+4, j, oakLeaves)
				}
			}
			// generate the top 2 layers of the tree's leaves (both 2x2 and one excluding corners)
			if i >= x-1 && i <= x+1 && j >= z-1 && j <= z+1 {
				c.SetBlock(i, surface+6, j, oakLeaves)

				if !isCorner(i, j, x-1, z-1, x+1, z+1) {
					c.SetBlock(i, surface+7, j, oakLeaves)
				}
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

func isCorner(x, z, minX, minZ, maxX, maxZ int32) bool {
	return (x == minX && z == minZ) || (x == maxX && z == maxZ) || (x == maxX && z == minZ) || (x == minZ && z == maxZ)
}

var grassBlock = block.GrassBlock{}
var dirt = block.Dirt{}
var bedrock = block.Bedrock{}
var oakLog = block.OakLog{Axis: block.AxisY}
var oakLeaves = block.OakLeaves{Distance: 1}
