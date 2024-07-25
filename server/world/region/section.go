package region

import (
	"math/bits"
)

// Section is a 16x16x16 chunk
type Section struct {
	blockLight, skyLight []byte
	y                    int8

	blockPalette []Block
	blockStates  []int64

	biomes struct {
		Data    []int64  `nbt:"data"`
		Palette []string `nbt:"palette"`
	} `nbt:"biomes"`

	blockBitsPerEntry int
}

func (s *Section) offset(x, y, z int) (long int, offset int) {
	usedBitsPerLong := (64 / s.blockBitsPerEntry) * s.blockBitsPerEntry
	blockNumber := (((y * 16) + z) * 16) + x
	startLong := (blockNumber * s.blockBitsPerEntry) / usedBitsPerLong
	stateOffset := (blockNumber * s.blockBitsPerEntry) % usedBitsPerLong

	return startLong, stateOffset
}

// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (sec *Section) block(x, y, z byte) Block {
	if len(sec.blockStates) == 0 {
		return sec.blockPalette[0]
	}

	long, off := sec.offset(int(x), int(y), int(z))
	l := sec.blockStates[long]
	index := (l >> off) & (1<<sec.blockBitsPerEntry - 1)

	return sec.blockPalette[index]
}

// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (sec *Section) setBlock(x, y, z byte, b Block) {
	state, ok := sec.index(b)
	if !ok {
		oldBPE := sec.blockBitsPerEntry
		sec.add(b)
		if oldBPE == sec.blockBitsPerEntry {
			state = int64(len(sec.blockPalette) - 1)
		} else {
			data := make([]int64, 4096/(64/sec.blockBitsPerEntry))
			newSec := Section{
				y:                 sec.y,
				blockLight:        sec.blockLight,
				skyLight:          sec.skyLight,
				blockPalette:      sec.blockPalette,
				blockStates:       data,
				biomes:            sec.biomes,
				blockBitsPerEntry: sec.blockBitsPerEntry,
			}
			for x := byte(0); x < 16; x++ {
				for y := byte(0); y < 16; y++ {
					for z := byte(0); z < 16; z++ {
						newSec.setBlock(x, y, z, sec.block(x, y, z))
					}
				}
			}
			*sec = newSec
		}
	}

	long, off := sec.offset(int(x), int(y), int(z))
	mask := int64(^((1<<sec.blockBitsPerEntry - 1) << off))

	sec.blockStates[long] &= mask
	sec.blockStates[long] |= state << off
}

func (sec *Section) index(b Block) (i int64, ok bool) {
	for i, entry := range sec.blockPalette {
		if entry.Name != b.Name {
			continue
		}
		if mapEqual(b.Properties, entry.Properties) {
			return int64(i), true
		}
	}
	return 0, false
}

func (sec *Section) add(b Block) {
	sec.blockPalette = append(sec.blockPalette, b)
	sec.blockBitsPerEntry = blockBitsPerEntry(len(sec.blockPalette))
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

func blockBitsPerEntry(paletteSize int) int {
	ln := bits.Len32(uint32(paletteSize) - 1)
	if ln <= 4 && ln != 0 {
		ln = 4
	}

	return ln
}
