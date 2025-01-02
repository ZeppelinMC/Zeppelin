package block

import (
	"strconv"
)

type CopperGrate struct {
	Waterlogged bool
}

func (b CopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CopperGrate) New(props BlockProperties) Block {
	return CopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}