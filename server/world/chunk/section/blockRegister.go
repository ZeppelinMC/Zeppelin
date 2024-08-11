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

//var blockBuf = bytes.NewReader(blockData)

func init() {
	/*rd := nbt.NewStaticReader(blockBuf)
	// Read the compound root type id and name
	_, _, _ = rd.ReadRoot(true)
	// reuse the compound reader struct
	var compoundReader nbt.CompoundReader

	for {
		// read a type id (Compound), name from the reader. The name is a block name in this example
		name, err, end := rd.Compound(&compoundReader)
		if end {
			break
		}
		if err != nil {
			return
		}

		// Read all the fields from this compound, "states" - list - func(int32,nbt.ListReader)
		if err := compoundReader.ReadAll(func(len int32, rd nbt.ListReader) {
			states := make([]blockState, len)
			for i := int32(0); i < len; i++ {
				states[i].Properties = make(map[string]string)
				// read a type id (compound) and read the specified values from it, Id: string, Properties: map[string]string
				rd.Read([]any{&states[i].Id, states[i].Properties})

			}
			blocks[name] = blockInfo{
				States: states,
			}
		}); err != nil {
			return
		}
	}*/
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
func BlockStateId(b Block) (id int32, ok bool) {
	name, props := b.Encode()
	block := blocks[name]

	for _, state := range block.States {
		if util.MapEqual(props, state.Properties) {
			return state.Id, true
		}
	}
	return 0, false
}
