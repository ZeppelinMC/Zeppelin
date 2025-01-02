package block

import (
	"strconv"
)

type WeatheredCopperGrate struct {
	Waterlogged bool
}

func (b WeatheredCopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:weathered_copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WeatheredCopperGrate) New(props BlockProperties) Block {
	return WeatheredCopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}