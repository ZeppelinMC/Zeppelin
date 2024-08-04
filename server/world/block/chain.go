package block

import (
	"strconv"
)

type Chain struct {
	Axis string
	Waterlogged bool
}

func (b Chain) Encode() (string, BlockProperties) {
	return "minecraft:chain", BlockProperties{
		"axis": b.Axis,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b Chain) New(props BlockProperties) Block {
	return Chain{
		Axis: props["axis"],
		Waterlogged: props["waterlogged"] != "false",
	}
}