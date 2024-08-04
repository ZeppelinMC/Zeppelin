package block

import (
	"strconv"
)

type DeepslateTileStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b DeepslateTileStairs) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_tile_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b DeepslateTileStairs) New(props BlockProperties) Block {
	return DeepslateTileStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}