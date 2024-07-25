package region

type Section struct {
	blockLight, skyLight []byte
	y                    int8

	blockPalette []anvilBlock
	blockStates  []int64

	biomes struct {
		Data    []int64  `nbt:"data"`
		Palette []string `nbt:"palette"`
	} `nbt:"biomes"`

	blockBitsPerEntry int
}

func (s Section) indexOffset(x, y, z int) (long int, offset int) {
	usedBitsPerLong := (64 / s.blockBitsPerEntry) * s.blockBitsPerEntry
	blockNumber := (((y * 16) + z) * 16) + x
	startLong := (blockNumber * s.blockBitsPerEntry) / usedBitsPerLong
	stateOffset := (blockNumber * s.blockBitsPerEntry) % usedBitsPerLong

	return startLong, stateOffset
}

// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (sec Section) block(x, y, z byte) anvilBlock {
	if sec.blockBitsPerEntry == 0 {
		return sec.blockPalette[0]
	}

	long, off := sec.indexOffset(int(x), int(y), int(z))
	l := sec.blockStates[long]
	index := (l >> off) & (1<<sec.blockBitsPerEntry - 1)

	return sec.blockPalette[index]
}

// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (sec Section) setBlock(x, y, z byte, b anvilBlock) {
	i, ok := sec.entryIndex(b)
	if !ok {
		panic("nah")
	}

	long, off := sec.indexOffset(int(x), int(y), int(z))
	l := sec.blockStates[long]
	pos := l & ((1<<sec.blockBitsPerEntry - 1) << off)

	l &= ^pos
	l |= int64(i) << pos

	sec.blockStates[long] = l
}

func (sec Section) entryIndex(b anvilBlock) (i int, ok bool) {
	for i, entry := range sec.blockPalette {
		if entry.Name != b.Name {
			continue
		}
		if mapEqual(b.Properties, entry.Properties) {
			return i, true
		}
	}
	return 0, false
}

func mapEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v1 := range a {
		v2, ok := b[k]
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}

	return true
}
