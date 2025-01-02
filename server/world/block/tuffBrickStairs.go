package block

import (
	"strconv"
)

type TuffBrickStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b TuffBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:tuff_brick_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b TuffBrickStairs) New(props BlockProperties) Block {
	return TuffBrickStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}