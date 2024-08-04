package block

import (
	"strconv"
)

type CutSandstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b CutSandstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:cut_sandstone_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b CutSandstoneSlab) New(props BlockProperties) Block {
	return CutSandstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}