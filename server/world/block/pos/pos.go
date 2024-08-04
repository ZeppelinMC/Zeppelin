package pos

func New(x, y, z int32) BlockPosition {
	return BlockPosition{x, y, z}
}

type BlockPosition [3]int32

// returns the absolute x position of the block in the world
func (b BlockPosition) X() int32 {
	return b[0]
}

// returns the absolute y position of the block in the world
func (b BlockPosition) Y() int32 {
	return b[1]
}

// returns the absolute z position of the block in the world
func (b BlockPosition) Z() int32 {
	return b[2]
}

// returns the chunk position x of this block
func (b BlockPosition) ChunkX() int32 {
	return b[0] >> 4
}

// returns the chunk position y of this block
func (b BlockPosition) Section() int32 {
	return b[1] >> 4
}

// returns the chunk position z of this block
func (b BlockPosition) ChunkZ() int32 {
	return b[2] >> 4
}

// returns the relative x position of the block in its section
func (b BlockPosition) SectionX() int32 {
	return b[0] & 0x0f
}

// returns the relative y position of the block in its section
func (b BlockPosition) SectionY() int32 {
	return b[1] & 0x0f
}

// returns the relative z position of the block in its section
func (b BlockPosition) SectionZ() int32 {
	return b[2] & 0x0f
}

func (b BlockPosition) Add(b1 BlockPosition) BlockPosition {
	return BlockPosition{
		b[0] + b1[0],
		b[1] + b1[1],
		b[2] + b1[2],
	}
}

func (b BlockPosition) Sub(b1 BlockPosition) BlockPosition {
	return BlockPosition{
		b[0] - b1[0],
		b[1] - b1[1],
		b[2] - b1[2],
	}
}
