package block

import (
	"strconv"
)

type BirchStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b BirchStairs) Encode() (string, BlockProperties) {
	return "minecraft:birch_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b BirchStairs) New(props BlockProperties) Block {
	return BirchStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}