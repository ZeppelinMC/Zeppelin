package blocks

import (
	_ "embed"
	"encoding/json"
	"reflect"
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

//go:embed blocks.json
var blockData []byte

func LoadBlockCache() error {
	return json.Unmarshal(blockData, &Blocks)
}
