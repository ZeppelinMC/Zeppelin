package chunk

import (
	"log"
	"math/bits"

	"github.com/dynamitemc/dynamite/server/block"
)

type section struct {
	data []int64

	bitsPerEntry int

	entries []block.Block

	ids []int32

	blockLight, skyLight []int8
}

func newSection(data []int64, blocks []blockEntry, bLight, sLight []int8) (s *section) {
	if len(blocks) == 0 {
		return nil
	}

	s = new(section)
	s.data = data

	s.entries = make([]block.Block, 0, len(blocks))
	for _, entry := range blocks {
		b := block.GetBlock(entry.Name)
		if entry.Properties != nil {
			b = b.New(entry.Properties)
		}

		id, ok := block.GetBlockId(b)
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

	s.blockLight = bLight
	s.skyLight = sLight
	return s
}
