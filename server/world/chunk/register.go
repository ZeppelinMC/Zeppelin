package chunk

import (
	_ "embed"
	"fmt"

	"github.com/aimjel/minecraft/nbt"
)

//go:embed blocks.nbt
var nbtBytes []byte

type blockInfo struct {
	Properties map[string][]string `json:"properties"`

	States []blockState
}

type blockState struct {
	Id         int  `json:"id"`
	Default    bool `json:"default"`
	Properties map[string]string
}

var blocks map[string]blockInfo

var registeredBlocks = map[string]Block{}

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

func GetBlock(name string) Block {
	if b, ok := registeredBlocks[name]; ok {
		return b
	}

	return NewUnknownBlock(name)
}

func DefaultBlock(name string) Block {
	b := GetBlock(name)
	if b == nil {
		return nil
	}
	block := blocks[b.EncodedName()]
	d, _ := defaultState(block.States)
	return b.New(name, d.Properties)
}

func defaultState(s []blockState) (_ blockState, ok bool) {
	for _, i := range s {
		if i.Default {
			return i, ok
		}
	}
	return
}

func GetBlockId(b Block) (int, bool) {
	block := blocks[b.EncodedName()]
	prop := b.Properties()
	if prop == nil {
		prop = map[string]string{}
	}

	for _, state := range block.States {
		if eq(state.Properties, prop) {
			return state.Id, true
		}
	}
	return 0, false
}

func eq(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}
