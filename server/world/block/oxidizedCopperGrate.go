package block

import (
	"strconv"
)

type OxidizedCopperGrate struct {
	Waterlogged bool
}

func (b OxidizedCopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OxidizedCopperGrate) New(props BlockProperties) Block {
	return OxidizedCopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}