package block

import (
	"strconv"
)

type WaxedCutCopperSlab struct {
	Waterlogged bool
	Type string
}

func (b WaxedCutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:waxed_cut_copper_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b WaxedCutCopperSlab) New(props BlockProperties) Block {
	return WaxedCutCopperSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}