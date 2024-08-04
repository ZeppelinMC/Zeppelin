package block

import (
	"strconv"
)

type SmoothStoneSlab struct {
	Type string
	Waterlogged bool
}

func (b SmoothStoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:smooth_stone_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b SmoothStoneSlab) New(props BlockProperties) Block {
	return SmoothStoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}