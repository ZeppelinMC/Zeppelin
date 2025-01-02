package block

import (
	"strconv"
)

type OxidizedCutCopperSlab struct {
	Type string
	Waterlogged bool
}

func (b OxidizedCutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_cut_copper_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OxidizedCutCopperSlab) New(props BlockProperties) Block {
	return OxidizedCutCopperSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}