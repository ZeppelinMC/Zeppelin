package block

import (
	"strconv"
)

type HangingRoots struct {
	Waterlogged bool
}

func (b HangingRoots) Encode() (string, BlockProperties) {
	return "minecraft:hanging_roots", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b HangingRoots) New(props BlockProperties) Block {
	return HangingRoots{
		Waterlogged: props["waterlogged"] != "false",
	}
}