// Package section provides methods of modifying chunk sections
package section

import (
	"fmt"
	"math"
	"math/bits"
	"reflect"
)

/* credit to https://github.com/aimjel for these calculations */

type AnvilBlock struct {
	Properties map[string]string
	Name       string
}

type AnvilBiomes struct {
	Data    []int64  `nbt:"data,omitempty"`
	Palette []string `nbt:"palette"`
}

type AnvilBlockStates struct {
	Data    []int64      `nbt:"data,omitempty"`
	Palette []AnvilBlock `nbt:"palette"`
}

// Section is a 16x16x16 chunk
type Section struct {
	BlockLight  []int8 `nbt:"BlockLight,omitempty"`
	SkyLight    []int8 `nbt:"SkyLight,omitempty"`
	Y           int8
	Biomes      AnvilBiomes      `nbt:"biomes"`
	BlockStates AnvilBlockStates `nbt:"block_states"`

	blockBitsPerEntry, biomeBitsPerEntry int
}

func (s *Section) Init() {
	s.blockBitsPerEntry = blockBitsPerEntry(len(s.BlockStates.Palette))
	s.biomeBitsPerEntry = biomeBitsPerEntry(len(s.Biomes.Palette))
}

func (s *Section) BPE() (block, biome int) {
	return s.blockBitsPerEntry, s.biomeBitsPerEntry
}

func (s *Section) offset(x, y, z int) (long, offset int) {
	blockNumber := ((((y * 16) + z) * 16) + x) * s.blockBitsPerEntry
	startLong := blockNumber >> 6
	stateOffset := blockNumber & 63

	return startLong, stateOffset
}

// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (s *Section) Block(x, y, z byte) AnvilBlock {
	if len(s.BlockStates.Data) == 0 {
		return s.BlockStates.Palette[0]
	}

	long, off := s.offset(int(x), int(y), int(z))
	l := s.BlockStates.Data[long]
	index := (l >> off) & (1<<s.blockBitsPerEntry - 1)

	return s.BlockStates.Palette[index]
}

func (s *Section) blockState(x, y, z byte) int64 {
	if len(s.BlockStates.Data) == 0 {
		return 0
	}

	long, off := s.offset(int(x), int(y), int(z))
	l := s.BlockStates.Data[long]
	index := (l >> off) & (1<<s.blockBitsPerEntry - 1)

	return index
}

// Sets the block at the position to the index in the palette. Errors if the state is bigger than the length of the palette or negative
// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (s *Section) setBlockState(x, y, z byte, state int64) error {
	if state < 0 || len(s.BlockStates.Data) <= int(state) {
		return fmt.Errorf("block state not in palette")
	}
	long, off := s.offset(int(x), int(y), int(z))
	mask := int64(^((1<<s.blockBitsPerEntry - 1) << off))

	s.BlockStates.Data[long] &= mask
	s.BlockStates.Data[long] |= state << off
	return nil
}

// Sets the block at the position and returns its new state (index in the block palette). Resizes the palette if needed
// X, Y, Z should be relative to the chunk section (AKA x&0x0f, y&0x0f, z&0x0f)
func (s *Section) SetBlock(x, y, z byte, b AnvilBlock) (state int64) {
	var ok bool
	state, ok = s.index(b)
	if !ok {
		panic("not now")
		oldBPE := s.blockBitsPerEntry
		s.add(b)
		state = int64(len(s.BlockStates.Palette) - 1)
		if oldBPE != s.blockBitsPerEntry {
			newDataLen := 4096 / (64 / s.blockBitsPerEntry)
			if len(s.BlockStates.Data) < newDataLen {
				s.BlockStates.Data = make([]int64, newDataLen)
			}
			clear(s.BlockStates.Data)

			for x := byte(0); x < 16; x++ {
				for y := byte(0); y < 16; y++ {
					for z := byte(0); z < 16; z++ {
						s.setBlockState(x, y, z, s.blockState(x, y, z))
					}
				}
			}
		}
	}

	long, off := s.offset(int(x), int(y), int(z))
	mask := int64(^((1<<s.blockBitsPerEntry - 1) << off))

	s.BlockStates.Data[long] &= mask
	s.BlockStates.Data[long] |= state << off

	return state
}

func (s *Section) index(b AnvilBlock) (i int64, ok bool) {
	for i, entry := range s.BlockStates.Palette {
		if reflect.DeepEqual(entry, b) { //todo optimize
			return int64(i), true
		}
	}
	return 0, false
}

func (s *Section) add(b AnvilBlock) {
	s.BlockStates.Palette = append(s.BlockStates.Palette, b)
	s.blockBitsPerEntry = blockBitsPerEntry(len(s.BlockStates.Palette))
}

func (s *Section) SkyLightLevel(x, y, z int) (byte, error) {
	index := (y << 8) | (z << 4) | x
	if index < 0 || int(index) > len(s.SkyLight) {
		return 0, fmt.Errorf("light index exceeds light length")
	}

	var mask uint8 = 0x0f
	var shift uint8 = 0
	if math.Remainder(float64(index), 2) != 0 {
		mask = 0xf0
		shift = 4
	}

	index /= 2

	return (uint8(s.SkyLight[index]) & mask) >> shift, nil
}

func (s *Section) SetSkylightLevel(x, y, z int, level byte) error {
	if level > 0x0F {
		return fmt.Errorf("light level must not exceed 4 bits (15)")
	}
	index := (y << 8) | (z << 4) | x
	if index < 0 || int(index) > len(s.SkyLight) {
		return fmt.Errorf("light index exceeds light length")
	}

	var mask uint8 = 0x0f
	if math.Remainder(float64(index), 2) != 0 {
		mask = 0xf0
		level <<= 4
	}

	index /= 2

	s.SkyLight[index] = s.SkyLight[index] & ^int8(mask)
	s.SkyLight[index] |= int8(level)

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
