package block

import (
	"strconv"
)

type Barrier struct {
	Waterlogged bool
}

func (b Barrier) Encode() (string, BlockProperties) {
	return "minecraft:barrier", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b Barrier) New(props BlockProperties) Block {
	return Barrier{
		Waterlogged: props["waterlogged"] != "false",
	}
}