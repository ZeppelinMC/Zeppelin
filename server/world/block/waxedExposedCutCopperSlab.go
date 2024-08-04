package block

import (
	"strconv"
)

type WaxedExposedCutCopperSlab struct {
	Type string
	Waterlogged bool
}

func (b WaxedExposedCutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_cut_copper_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedExposedCutCopperSlab) New(props BlockProperties) Block {
	return WaxedExposedCutCopperSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}