package block

import (
	"strconv"
)

type HornCoral struct {
	Waterlogged bool
}

func (b HornCoral) Encode() (string, BlockProperties) {
	return "minecraft:horn_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b HornCoral) New(props BlockProperties) Block {
	return HornCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}