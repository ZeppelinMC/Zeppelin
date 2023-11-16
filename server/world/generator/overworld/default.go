package overworld

import (
	"io"

	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/block"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type DefaultGenerator struct {
	g bool
}

func (f *DefaultGenerator) GenerateChunk(cx, cz int32) (*chunk.Chunk, error) {
	if f.g {
		return nil, io.EOF
	}
	f.g = true
	logger.Println("generating chunk ", cx, cz)
	c := chunk.Chunk{}
	c.SetPosition(cx, cz)
	/*for x := int64(cx * 16); x < int64(cx*16)+16; x++ {
		for z := int64(cz * 16); z < int64(cz*16)+16; z++ {
			c.SetBlock(x, 0, z, block.Bedrock{})
			c.SetBlock(x, 1, z, block.Bedrock{})
			c.SetBlock(x, 2, z, block.Bedrock{})
			c.SetBlock(x, 3, z, block.Bedrock{})
		}
	}*/
	for x := int64(0); x < 16; x++ {
		for z := int64(0); z < 16; z++ {
			c.SetBlock(x, -64, z, block.Bedrock{})
			c.SetBlock(x, -63, z, block.Dirt{})
			c.SetBlock(x, -62, z, block.Dirt{})
			c.SetBlock(x, -61, z, block.GrassBlock{})
		}
	}
	return &c, nil
}
