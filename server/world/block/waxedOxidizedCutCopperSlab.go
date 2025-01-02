package block

import (
	"strconv"
)

type WaxedOxidizedCutCopperSlab struct {
	Type string
	Waterlogged bool
}

func (b WaxedOxidizedCutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_cut_copper_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedOxidizedCutCopperSlab) New(props BlockProperties) Block {
	return WaxedOxidizedCutCopperSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}