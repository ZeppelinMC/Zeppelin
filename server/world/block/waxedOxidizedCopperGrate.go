package block

import (
	"strconv"
)

type WaxedOxidizedCopperGrate struct {
	Waterlogged bool
}

func (b WaxedOxidizedCopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedOxidizedCopperGrate) New(props BlockProperties) Block {
	return WaxedOxidizedCopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}