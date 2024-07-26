package terrain

import "github.com/zeppelinmc/zeppelin/server/world/region"

// A superflat chunk generator. A superflat world is just 4 layers. One bedrock, two dirt, and one grass block, so it is very easy to implement
type SuperflatTerrain struct {
}

func (SuperflatTerrain) NewChunk(cx, cz int32) region.Chunk {
	c := region.NewChunk(cx, cz)

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
