package block

import (
	"strconv"
)

type WaxedWeatheredCutCopperSlab struct {
	Type string
	Waterlogged bool
}

func (b WaxedWeatheredCutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_cut_copper_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedWeatheredCutCopperSlab) New(props BlockProperties) Block {
	return WaxedWeatheredCutCopperSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}