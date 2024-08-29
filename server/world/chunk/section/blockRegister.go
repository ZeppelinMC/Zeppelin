package section

import (
	_ "embed"
	"io"

	"github.com/zeppelinmc/zeppelin/server/world/block/blockstates"
	"github.com/zeppelinmc/zeppelin/util"
)

var blocks = make(map[string]blockstates.Block)

type RReaderAt interface {
	io.Reader
	io.ReaderAt
}

var header map[string]blockstates.BlockLocation
var blockFile RReaderAt

func ImportStates(f RReaderAt) (err error) {
	blockFile = f
	header, err = blockstates.ReadHeader(blockFile)

	return
}

var registeredBlocks = make(map[string]Block)

// Registers a block struct that will be used for creating blocks with the name returned by the block's Encode function
func RegisterBlock(b Block) {
	name, _ := b.Encode()
	registeredBlocks[name] = b
}

// Returns the block struct found for the block name
func GetBlock(name string) Block {
	if b, ok := registeredBlocks[name]; ok {
		return b
	}
	return UnknownBlock{name: name}
}

// returns the block state id for this block
func BlockStateId(b Block) (id int32, ok bool) {
	name, props := b.Encode()
	block, ok := blocks[name]

	if !ok {
		block, _ = blockstates.ReadBlock(blockFile, header[name])
		blocks[name] = block
	}

	for _, state := range block {
		if util.MapEqual(props, state.Properties) {
			return state.Id, true
		}
	}
	return 0, false
}
