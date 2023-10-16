package block

import (
	_ "embed"
	"fmt"
	"github.com/aimjel/minecraft/nbt"
)

//go:embed blocks.nbt
var nbtBytes []byte

type blockInfo struct {
	Properties map[string][]string `json:"properties"`

	States []struct {
		Id         int  `json:"id"`
		Default    bool `json:"default"`
		Properties map[string]string
	}
}

var blocks map[string]blockInfo

var registeredBlocks map[string]Block

func init() {
	if err := nbt.Unmarshal(nbtBytes, &blocks); err != nil {
		panic(err)
	}
}

func RegisterBlock(b Block) error {
	_, ok := registeredBlocks[b.EncodedName()]

	if ok {
		return fmt.Errorf("%v block is already registered", b.EncodedName())
	}

	registeredBlocks[b.EncodedName()] = b
	return nil
}
