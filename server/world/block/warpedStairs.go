package block

import (
	"strconv"
)

type WarpedStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b WarpedStairs) Encode() (string, BlockProperties) {
	return "minecraft:warped_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b WarpedStairs) New(props BlockProperties) Block {
	return WarpedStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}