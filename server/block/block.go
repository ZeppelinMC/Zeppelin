package block

type Block interface {
	EncodedName() string

	New(map[string]string) Block

	Properties() map[string]string
}

func GetBlock(name string) Block {
	if b, ok := registeredBlocks[name]; ok {
		return b
	}

	return &UnknownBlock{encodedName: name}
}

func GetBlockId(b Block) (int, bool) {
	block := blocks[b.EncodedName()]

	for _, state := range block.States {

		if eq(state.Properties, b.Properties()) {
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
