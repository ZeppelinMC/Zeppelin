package chunk

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/heightmaps"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
)

const MinChunkY = -4

type BlockEntity struct {
	Id         string `nbt:"id"`
	X          int32  `nbt:"x"`
	Y          int32  `nbt:"y"`
	Z          int32  `nbt:"z"`
	KeepPacked bool   `nbt:"keepPacked"`

	// for chest entities
	Items container.Container `nbt:"Items,omitempty"`
}

type Chunk struct {
	X, Y, Z    int32
	Heightmaps heightmaps.Heightmaps

	Sections      []*section.Section
	BlockEntities []BlockEntity
}

func NewChunk(x, z int32) Chunk {
	var airBlock = section.GetBlock("minecraft:air")
	c := Chunk{
		Y: MinChunkY,
		X: x,
		Z: z,

		Sections: make([]*section.Section, 24),
	}

	for i := range c.Sections {
		c.Sections[i] = section.New(
			int8(i-MinChunkY),
			[]section.Block{airBlock},
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
func (c *Chunk) Block(x, y, z int32) (section.Block, error) {
	secIndex := (y >> 4) - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.Sections)) {
		return nil, fmt.Errorf("null section")
	}
	sec := c.Sections[secIndex]

	return sec.Block(byte(x), byte(y)&0x0f, byte(z)), nil
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetBlock(x, y, z int32, b section.Block) (state int64, err error) {
	secIndex := (y >> 4) - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.Sections)) {
		return 0, fmt.Errorf("null section")
	}
	sec := c.Sections[secIndex]
	return sec.SetBlock(byte(x), byte(y)&0x0f, byte(z), b), nil
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetBlockState(x, y, z int32, state int64) error {
	secIndex := (y >> 4) - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.Sections)) {
		return fmt.Errorf("null section")
	}
	sec := c.Sections[secIndex]
	return sec.SetBlockState(byte(x), byte(y)&0x0f, byte(z), state)
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetSkylightLevel(x, y, z int32, value byte) error {
	secIndex := (y >> 4) - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.Sections)) {
		return fmt.Errorf("null section")
	}
	sec := c.Sections[secIndex]
	return sec.SetSkylightLevel(int(x), int(y)&0x0f, int(z), value)
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetBlockLightLevel(x, y, z int32, value byte) error {
	secIndex := (y >> 4) - c.Y
	if secIndex < 0 || secIndex >= int32(len(c.Sections)) {
		return fmt.Errorf("null section")
	}
	sec := c.Sections[secIndex]
	return sec.SetBlocklightLevel(int(x), int(y)&0x0f, int(z), value)
}

// Returns the block entity at the position. All of the position values should be absolute (aka (chunkPos<<4)+pos)
func (c *Chunk) BlockEntity(x, y, z int32) (*BlockEntity, bool) {
	for _, entity := range c.BlockEntities {
		if entity.X == x && entity.Y == y && entity.Z == z {
			return &entity, true
		}
	}
	return nil, false
}
