package block

import (
	"strconv"
)

type SmoothSandstoneStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b SmoothSandstoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:smooth_sandstone_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SmoothSandstoneStairs) New(props BlockProperties) Block {
	return SmoothSandstoneStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}