package section

import (
	_ "embed"
	"github.com/aimjel/minecraft/nbt"
	"github.com/zeppelinmc/zeppelin/util"
)

type blockState struct {
	Id         int32             `json:"id"`
	Properties map[string]string `json:"properties"`
}

type blockInfo struct {
	States []blockState `json:"states"`
}

var blocks = make(map[string]blockInfo)

//go:embed data/blocks.nbt
var blockData []byte

func init() {
	nbt.Unmarshal(blockData, &blocks)
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
func BlockStateId(b AnvilBlock) (id int32, ok bool) {
	block := blocks[b.Name]

	for _, state := range block.States {
		if util.MapEqual(b.Properties, state.Properties) {
			return state.Id, true
		}
	}
	return 0, false
}
