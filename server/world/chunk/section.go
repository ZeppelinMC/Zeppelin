package chunk

import (
	"github.com/aimjel/nitrate/server/block"
	"log"
	"math"
	"math/bits"
)

// Section represents a 16x16x16.
type Section struct {
	bitsPerEntry int

	ids []int32

	data []int64

	blocks []block.Block

	skyLight []int8

	blockLight []int8
}

func NewSection(blocks []block.Block, data []int64, skyLight, blockLight []int8) *Section {
	s := &Section{
		data:       data,
		ids:        make([]int32, 0, len(blocks)),
		blocks:     blocks,
		skyLight:   skyLight,
		blockLight: blockLight,
	}
	for _, b := range blocks {
		id, ok := block.GetBlockId(b)
		if !ok {
			log.Panicf("unable to find block id for %+v\n", b)
		}

		s.ids = append(s.ids, int32(id))
	}

	ln := bits.Len32(uint32(len(blocks)))
	if ln == 0 {
		s.bitsPerEntry = 0
		s.blocks = append(s.blocks, block.Air{})
		id, _ := block.GetBlockId(block.Air{})
		s.ids = append(s.ids, int32(id))
		return s
	}
	if ln == 1 {
		//use indirect palette
		s.bitsPerEntry = 0
		return s
	}

	if ln <= 4 {
		s.bitsPerEntry = ln
	}

	if data == nil {
		s.data = make([]int64, 4096/(64/ln))
	}
	return s
}

func (s *Section) indexOffset(x, y, z int) (int, int) {
	usedBitsPerLong := (64 / s.bitsPerEntry) * s.bitsPerEntry
	blockNumber := (((y * 16) + z) * 16) + x
	startLong := (blockNumber * s.bitsPerEntry) / usedBitsPerLong
	stateOffset := (blockNumber * s.bitsPerEntry) % usedBitsPerLong

	return startLong, stateOffset
}

func (s *Section) GetBlockAt(x, y, z int) block.Block {
	if s.bitsPerEntry == 0 {
		return s.blocks[0]
	}

	i, offset := s.indexOffset(x, y, z)

	states := s.data[i]
	states >>= offset

	data := states & (1<<s.bitsPerEntry - 1)

	return s.blocks[data]
}

func (s *Section) SetBlock(x, y, z int, b block.Block) {
	state, ok := s.index(b)
	if !ok {
		s.addBlock(b)
		state = int64(len(s.ids)) - 1
	}

	i, offset := s.indexOffset(x, y, z)
	mask := int64(^((1<<s.bitsPerEntry - 1) << offset))
	s.data[i] &= mask
	s.data[i] |= state << offset
}

func (s *Section) index(b block.Block) (int64, bool) {
	for k, v := range s.blocks {
		if v == b {
			return int64(k), true
		}
	}

	return 0, false
}

// addBlock adds the block to the palette
func (s *Section) addBlock(b block.Block) {
	id, ok := block.GetBlockId(b)
	if !ok {
		panic("couldnt find block id")
	}

	s.ids = append(s.ids, int32(id))
	s.blocks = append(s.blocks, b)

	ln := bits.Len32(uint32(len(s.blocks)) - 1)
	if ln <= 4 && ln != 0 {
		ln = 4
	}

	if ln == 0 {
		s.bitsPerEntry = ln
		return
	}

	if ln != s.bitsPerEntry {

		data := make([]int64, 4096/(64/ln))
		newSection := Section{
			bitsPerEntry: ln,
			ids:          s.ids,
			data:         data,
			blocks:       s.blocks,
			skyLight:     s.skyLight,
			blockLight:   s.blockLight,
		}

		for x := 0; x < 16; x++ {
			for y := 0; y < 16; y++ {
				for z := 0; z < 16; z++ {
					newSection.SetBlock(x, y, z, s.GetBlockAt(x, y, z))
				}
			}
		}

		*s = newSection
	}
}

func (s *Section) setSkyLight(x, y, z int, level uint8) {
	if s.skyLight == nil {
		s.skyLight = make([]int8, 2048)
	}

	index := (y << 8) | (z << 4) | x

	remain := math.Remainder(float64(index), 2)

	var mask uint8 = 0x0f
	if remain != 0 {
		mask = 0xf0
		level <<= 4
	}

	index /= 2

	s.skyLight[index] &= ^int8(mask)

	s.skyLight[index] |= int8(level)
}
