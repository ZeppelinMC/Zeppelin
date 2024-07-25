package region

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

	sections []*Section
}

func NewChunk(x, z int32) Chunk {
	c := Chunk{
		Y: MinChunkY,
		X: x,
		Z: z,

		sections: make([]*Section, 24),
	}
	for i := range c.sections {
		s := Section{}
		s.y = int8(i - MinChunkY)
		s.biomes.Palette = []string{"minecraft:plains"}
		s.blockPalette = []Block{{Name: "minecraft:air"}}
		s.skyLight = fullLightBuffer

		c.sections[i] = &s
	}

	return c
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) Block(x, y, z int32) Block {
	sec := c.sections[y/16-c.Y]

	return sec.block(byte(x), byte(y)&0x0f, byte(z))
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetBlock(x, y, z int32, b Block) {
	sec := c.sections[y/16-c.Y]
	sec.setBlock(byte(x), byte(y)&0x0f, byte(z), b)
}
