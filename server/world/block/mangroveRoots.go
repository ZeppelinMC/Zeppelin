package block

import (
	"strconv"
)

type MangroveRoots struct {
	Waterlogged bool
}

func (b MangroveRoots) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_roots", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MangroveRoots) New(props BlockProperties) Block {
	return MangroveRoots{
		Waterlogged: props["waterlogged"] != "false",
	}
}