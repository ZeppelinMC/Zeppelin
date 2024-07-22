package blocks

import (
	_ "embed"
	"reflect"

	"github.com/zeppelinmc/zeppelin/nbt"
)

type Block struct {
	States []struct {
		Id         int32             `json:"id"`
		Properties map[string]string `json:"properties"`
	} `json:"states"`
}

func (b Block) FindState(properties map[string]string) (id int32, ok bool) {
	for _, state := range b.States {
		if reflect.DeepEqual(state.Properties, properties) {
			return state.Id, true
		}
	}
	return 0, false
}

var Blocks map[string]Block

//go:embed blocks.nbt
var blockData []byte

func LoadBlockCache() error {
	_, err := nbt.Unmarshal(blockData, &Blocks)
	return err
}
