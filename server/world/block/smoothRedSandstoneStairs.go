package block

import (
	"strconv"
)

type SmoothRedSandstoneStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b SmoothRedSandstoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:smooth_red_sandstone_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SmoothRedSandstoneStairs) New(props BlockProperties) Block {
	return SmoothRedSandstoneStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}