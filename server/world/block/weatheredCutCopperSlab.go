package block

import (
	"strconv"
)

type WeatheredCutCopperSlab struct {
	Type string
	Waterlogged bool
}

func (b WeatheredCutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:weathered_cut_copper_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WeatheredCutCopperSlab) New(props BlockProperties) Block {
	return WeatheredCutCopperSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}