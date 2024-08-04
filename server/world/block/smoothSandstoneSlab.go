package block

import (
	"strconv"
)

type SmoothSandstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b SmoothSandstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:smooth_sandstone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SmoothSandstoneSlab) New(props BlockProperties) Block {
	return SmoothSandstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}