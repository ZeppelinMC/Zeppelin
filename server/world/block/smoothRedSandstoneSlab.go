package block

import (
	"strconv"
)

type SmoothRedSandstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b SmoothRedSandstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:smooth_red_sandstone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SmoothRedSandstoneSlab) New(props BlockProperties) Block {
	return SmoothRedSandstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}