package block

import (
	"strconv"
)

type BrickStairs struct {
	Shape string
	Waterlogged bool
	Facing string
	Half string
}

func (b BrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:brick_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b BrickStairs) New(props BlockProperties) Block {
	return BrickStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}