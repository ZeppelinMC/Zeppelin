package block

import (
	"strconv"
)

type CutRedSandstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b CutRedSandstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:cut_red_sandstone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CutRedSandstoneSlab) New(props BlockProperties) Block {
	return CutRedSandstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}