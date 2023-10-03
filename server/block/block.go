package block

import "reflect"

type Block interface {
	EncodedName() string

	New(map[string]string) Block

	Properties() map[string]string
}

func GetBlock(name string) Block {
	return &UnknownBlock{encodedName: name}
}

func GetBlockId(b Block) (int, bool) {
	block := jsonBlocks[b.EncodedName()]

	for _, state := range block.States {
		if reflect.DeepEqual(state.Properties, b.Properties()) {
			return state.Id, true
		}
	}
	return 0, false
}

type UnknownBlock struct {
	encodedName string
	properties  map[string]string
}

func (u UnknownBlock) EncodedName() string {
	return u.encodedName
}

func (u UnknownBlock) New(m map[string]string) Block {
	return UnknownBlock{encodedName: u.encodedName, properties: m}
}

func (u UnknownBlock) Properties() map[string]string {
	return u.properties
}
