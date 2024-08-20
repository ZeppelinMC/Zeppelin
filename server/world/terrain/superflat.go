package terrain

import (
	"math/rand"

	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

// superflat chunk generator
type SuperflatTerrain struct {
}

func (SuperflatTerrain) NewChunk(cx, cz int32) chunk.Chunk {
	c := chunk.NewChunk(cx, cz)

	for x := int32(0); x < 16; x++ {
		for z := int32(0); z < 16; z++ {
			c.SetBlock(x, 4, z, grassBlock)
			c.SetBlock(x, 2, z, dirt)
			c.SetBlock(x, 1, z, dirt)
			c.SetBlock(x, 0, z, bedrock)
		}
	}

	return c
}

func (SuperflatTerrain) GenerateWorldSpawn() (x, y, z int32) {
	return rand.Int31n(160), 5, rand.Int31n(160)
}
