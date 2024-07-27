package region

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/server/world/region/block"
	"github.com/zeppelinmc/zeppelin/server/world/region/section"
)

type Generator interface {
	NewChunk(x, z int32) Chunk
}

const MinChunkY = -4

func init() {
	for i := range fullLightBuffer {
		fullLightBuffer[i] = 0xFF
	}
}

type Chunk struct {
	X, Y, Z    int32
	Heightmaps Heightmaps

	sections []*section.Section
}

func NewChunk(x, z int32) Chunk {
	c := Chunk{
		Y: MinChunkY,
		X: x,
		Z: z,

		sections: make([]*section.Section, 24),
	}

	for i := range c.sections {
		c.sections[i] = section.New(
			int8(i-MinChunkY),
			[]block.Block{{Name: "minecraft:air"}},
			nil,
			[]string{"minecraft:plains"},
			nil,
			fullLightBuffer,
			nil,
		)
	}

	return c
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) Block(x, y, z int32) (block.Block, error) {
	secIndex := y/16 - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.sections)) {
		return block.Block{}, fmt.Errorf("null section")
	}
	sec := c.sections[secIndex]

	return sec.Block(byte(x), byte(y)&0x0f, byte(z)), nil
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetBlock(x, y, z int32, b block.Block) (state int64, err error) {
	secIndex := y/16 - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.sections)) {
		return 0, fmt.Errorf("null section")
	}
	sec := c.sections[secIndex]
	return sec.SetBlock(byte(x), byte(y)&0x0f, byte(z), b), nil
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetBlockState(x, y, z int32, state int64) error {
	secIndex := y/16 - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.sections)) {
		return fmt.Errorf("null section")
	}
	sec := c.sections[secIndex]
	return sec.SetBlockState(byte(x), byte(y)&0x0f, byte(z), state)
}
