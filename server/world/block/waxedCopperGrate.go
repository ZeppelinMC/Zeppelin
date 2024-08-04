package block

import (
	"strconv"
)

type WaxedCopperGrate struct {
	Waterlogged bool
}

func (b WaxedCopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:waxed_copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedCopperGrate) New(props BlockProperties) Block {
	return WaxedCopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}