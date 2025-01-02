package block

import (
	"strconv"
)

type WaxedWeatheredCopperGrate struct {
	Waterlogged bool
}

func (b WaxedWeatheredCopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedWeatheredCopperGrate) New(props BlockProperties) Block {
	return WaxedWeatheredCopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}