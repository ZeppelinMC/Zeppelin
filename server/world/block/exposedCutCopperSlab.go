package block

import (
	"strconv"
)

type ExposedCutCopperSlab struct {
	Type string
	Waterlogged bool
}

func (b ExposedCutCopperSlab) Encode() (string, BlockProperties) {
	return "minecraft:exposed_cut_copper_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b ExposedCutCopperSlab) New(props BlockProperties) Block {
	return ExposedCutCopperSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}