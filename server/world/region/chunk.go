package region

func init() {
	for i := range fullLightBuffer {
		fullLightBuffer[i] = 0xFF
	}
}

type Chunk struct {
	X, Y, Z    int32
	Heightmaps Heightmaps

	sections []Section
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) Block(x, y, z int32) anvilBlock {
	sec := c.sections[y/16-c.Y]

	return sec.block(byte(x), byte(y)&0x0f, byte(z))
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (c *Chunk) SetBlock(x, y, z int32, b anvilBlock) {

	sec := c.sections[y/16-c.Y]
	sec.setBlock(byte(x), byte(y)&0x0f, byte(z), b)
}
