package block

import (
	"strconv"
)

type CutCopperSlab struct {
	Type string
	Waterlogged bool
}

func (b CutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:cut_copper_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b CutCopperSlab) New(props BlockProperties) Block {
	return CutCopperSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}