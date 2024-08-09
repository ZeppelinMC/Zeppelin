// Package section provides methods of modifying chunk sections
package section

import (
	"fmt"
	"math"
	"math/bits"
)

/* credit to https://github.com/aimjel for these calculations */

func New(y int8, blockPalette []Block, blockStates []int64, biomePalette []string, biomesData []int64, skylight, blocklight []byte) *Section {
	s := &Section{
		y:                 y,
		blockPalette:      blockPalette,
		blockStates:       blockStates,
		skyLight:          skylight,
		blockLight:        blocklight,
		blockBitsPerEntry: blockBitsPerEntry(len(blockPalette)),
		biomeBitsPerEntry: biomeBitsPerEntry(len(biomePalette)),
	}
	s.biomes.Palette = biomePalette
	s.biomes.Data = biomesData
	return s
}

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
	biomeBitsPerEntry int
}

// Returns the lighting data for the section
func (sec *Section) Light() (skyLight, blockLight []byte) {
	return sec.skyLight, sec.blockLight
}

// Returns the block states data for this section
func (sec *Section) BlockStates() (bpe int, palette []Block, states []int64) {
	return sec.blockBitsPerEntry, sec.blockPalette, sec.blockStates
}

// Returns the biome data for this section
func (sec *Section) Biomes() (bpe int, palette []string, states []int64) {
	return sec.biomeBitsPerEntry, sec.biomes.Palette, sec.biomes.Data
}

func (sec *Section) Y() int8 {
	return sec.y
}

func (s *Section) offset(x, y, z int) (long int, offset int) {
	blockNumber := (((y * 16) + z) * 16) + x
	startLong := (blockNumber * s.blockBitsPerEntry) >> 6
	stateOffset := (blockNumber * s.blockBitsPerEntry) & 63

	return startLong, stateOffset
}

// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (sec *Section) Block(x, y, z byte) Block {
	if len(sec.blockStates) == 0 {
		return sec.blockPalette[0]
	}

	long, off := sec.offset(int(x), int(y), int(z))
	l := sec.blockStates[long]
	index := (l >> off) & (1<<sec.blockBitsPerEntry - 1)

	return sec.blockPalette[index]
}

// Sets the block at the position to the index in the palette. Errors if the state is bigger than the length of the palette or negative
// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (sec *Section) SetBlockState(x, y, z byte, state int64) error {
	if state < 0 || len(sec.blockStates) <= int(state) {
		return fmt.Errorf("block state not in palette")
	}
	long, off := sec.offset(int(x), int(y), int(z))
	mask := int64(^((1<<sec.blockBitsPerEntry - 1) << off))

	sec.blockStates[long] &= mask
	sec.blockStates[long] |= state << off
	return nil
}

// Sets the block at the position and returns its new state (index in the block palette). Resizes the palette if needed
// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (sec *Section) SetBlock(x, y, z byte, b Block) (state int64) {
	state, ok := sec.index(b)
	if !ok {
		oldBPE := sec.blockBitsPerEntry
		sec.add(b)
		state = int64(len(sec.blockPalette) - 1)
		if oldBPE != sec.blockBitsPerEntry {
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
						newSec.SetBlock(x, y, z, sec.Block(x, y, z))
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

	return state
}

func (sec *Section) index(b Block) (i int64, ok bool) {
	for i, entry := range sec.blockPalette {
		if entry == b {
			return int64(i), true
		}
	}
	return 0, false
}

func (sec *Section) add(b Block) {
	sec.blockPalette = append(sec.blockPalette, b)
	sec.blockBitsPerEntry = blockBitsPerEntry(len(sec.blockPalette))
}

func (sec *Section) SkyLightLevel(x, y, z int) (byte, error) {
	index := (y << 8) | (z << 4) | x
	if index < 0 || int(index) > len(sec.skyLight) {
		return 0, fmt.Errorf("light index exceeds light length")
	}

	var mask uint8 = 0x0f
	var shift uint8 = 0
	if math.Remainder(float64(index), 2) != 0 {
		mask = 0xf0
		shift = 4
	}

	index /= 2

	return (sec.skyLight[index] & mask) >> shift, nil
}

func (sec *Section) SetSkylightLevel(x, y, z int, level byte) error {
	if level > 0x0F {
		return fmt.Errorf("light level must not exceed 4 bits (15)")
	}
	index := (y << 8) | (z << 4) | x
	if index < 0 || int(index) > len(sec.skyLight) {
		return fmt.Errorf("light index exceeds light length")
	}

	var mask uint8 = 0x0f
	if math.Remainder(float64(index), 2) != 0 {
		mask = 0xf0
		level <<= 4
	}

	index /= 2

	sec.skyLight[index] &= ^mask
	sec.skyLight[index] |= level

	return nil
}

func (sec *Section) BlockLightLevel(x, y, z int) (byte, error) {
	index := (y << 8) | (z << 4) | x
	if index < 0 || int(index) > len(sec.blockLight) {
		return 0, fmt.Errorf("light index exceeds light length")
	}

	var mask uint8 = 0x0f
	var shift uint8 = 0
	if math.Remainder(float64(index), 2) != 0 {
		mask = 0xf0
		shift = 4
	}

	index /= 2

	return (sec.blockLight[index] & mask) >> shift, nil
}

func (sec *Section) SetBlocklightLevel(x, y, z int, level byte) error {
	if level > 0x0F {
		return fmt.Errorf("light level must not exceed 4 bits (15)")
	}
	index := (y << 8) | (z << 4) | x
	if index < 0 || int(index) > len(sec.blockLight) {
		return fmt.Errorf("light index exceeds light length")
	}

	var mask uint8 = 0x0f
	if math.Remainder(float64(index), 2) != 0 {
		mask = 0xf0
		level <<= 4
	}

	index /= 2

	sec.blockLight[index] &= ^mask
	sec.blockLight[index] |= level

	return nil
}

func blockBitsPerEntry(paletteSize int) int {
	ln := bits.Len32(uint32(paletteSize) - 1)
	if ln <= 4 && ln != 0 {
		ln = 4
	}

	//log.Println(paletteSize, ln)

	return ln
}

func biomeBitsPerEntry(paletteSize int) int {
	return bits.Len32(uint32(paletteSize - 1))
}
