package block

import (
	"strconv"
)

type WaxedExposedCopperGrate struct {
	Waterlogged bool
}

func (b WaxedExposedCopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedExposedCopperGrate) New(props BlockProperties) Block {
	return WaxedExposedCopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}