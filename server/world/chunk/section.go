package chunk

import (
	"fmt"
	"log"
	"math"
	"math/bits"

	"github.com/dynamitemc/dynamite/logger"
)

type section struct {
	data []int64

	bitsPerEntry int

	entries []Block

	ids []int32

	blockLight, skyLight []int8
}

func newSection(data []int64, blocks []blockEntry, bLight, sLight []int8) (s *section) {
	if len(blocks) == 0 {
		return nil
	}

	s = new(section)
	s.data = data

	s.entries = make([]Block, 0, len(blocks))
	s.ids = make([]int32, 0, len(blocks))
	for _, entry := range blocks {
		b := GetBlock(entry.Name)
		if entry.Properties != nil {
			b = b.New(entry.Name, entry.Properties)
		}

		id, ok := GetBlockId(b)
		if !ok {
			log.Panicf("unable to find block id for %+v\n", b)
		}

		s.ids = append(s.ids, int32(id))
		s.entries = append(s.entries, b)
	}

	ln := bits.Len32(uint32(len(blocks)) - 1)
	if ln <= 4 && ln != 0 {
		ln = 4
	}

	s.bitsPerEntry = ln

	//if s.bitsPerEntry == 0 {
	//	panic("shit!!!!!!!!!!!!!!!!")
	//}

	s.blockLight = bLight
	s.skyLight = sLight
	return s
}

// it be not workin!!
//this is fine
//the problem is a section has a bits per entry with zero
// cant do 64 / 0

//but its impossible for that to happen so i need more info
// idk how to give you more info
//debug the chunk which has a bits per entry with zero?

// lets see
// open terminal

// indexOffset returns which index the xyz are in,
// and the offset within the 64-bit value.
func (s *section) indexOffset(x, y, z int) (int, int) {
	usedBitsPerLong := (64 / s.bitsPerEntry) * s.bitsPerEntry
	blockNumber := (((y * 16) + z) * 16) + x
	startLong := (blockNumber * s.bitsPerEntry) / usedBitsPerLong
	stateOffset := (blockNumber * s.bitsPerEntry) % usedBitsPerLong

	return startLong, stateOffset
}
func (s *section) GetBlockAt(x, y, z int) Block {
	if s.bitsPerEntry == 0 {
		return s.entries[0]
	}

	i, offset := s.indexOffset(x, y, z)

	states := s.data[i]
	states >>= offset

	data := states & (1<<s.bitsPerEntry - 1)

	return s.entries[data]
}

func (s *section) setBlockAt(x, y, z int, b Block) {
	newState, ok := s.index(b)
	if !ok {
		old := s.bitsPerEntry
		fmt.Printf("adding %#v to palette entries\n\r", b)
		s.addEntry(b)

		if s.bitsPerEntry != old {
			logger.Println("RESIZING STATES SLICE")
			//the amount of bits we need for a new chunk
			newBits := (16 * 16 * 16) * s.bitsPerEntry

			//gets the amount of bits we can use in a 64-bit type
			newUsedBitsPerLong := (64 / s.bitsPerEntry) * s.bitsPerEntry

			newSec := section{
				entries: s.entries,
				ids:     s.ids,

				bitsPerEntry: s.bitsPerEntry,
				//create a new slice which can hold all the blocks
				data: make([]int64, int(math.Ceil(float64(newBits/newUsedBitsPerLong)))),

				skyLight:   s.skyLight,
				blockLight: s.blockLight,
			}

			s.bitsPerEntry = old
			for i := 0; i < 16; i++ {
				for j := 0; j < 16; j++ {
					for k := 0; k < 16; k++ {
						newSec.setBlockAt(i, j, k, s.GetBlockAt(i, j, k))
					}
				}
			}

			*s = newSec
			newState, _ = s.index(b)
		}
	}

	i, offset := s.indexOffset(x, y, z)
	mask := int64(^((1<<s.bitsPerEntry - 1) << offset))
	s.data[i] &= mask
	s.data[i] |= newState << offset
}

func (s *section) addEntry(b Block) {
	s.entries = append(s.entries, b)

	bitsPerEntry := bits.Len(uint(len(s.entries) - 1))
	if bitsPerEntry < 4 {
		bitsPerEntry = 4
	}

	s.bitsPerEntry = bitsPerEntry

	id, ok := GetBlockId(b)
	if !ok {
		log.Panicf("unable to find block id for %+v\n", b)
	}

	s.ids = append(s.ids, int32(id))
}

func (s *section) index(b Block) (int64, bool) {
	for k, v := range s.entries {
		//temp
		if v.EncodedName() == b.EncodedName() {
			return int64(k), true
		}
	}

	return 0, false
}

func (s *section) Blocks() []Block {
	return s.entries
}
