package block

import (
	"strconv"
)

type SmoothQuartzStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b SmoothQuartzStairs) Encode() (string, BlockProperties) {
	return "minecraft:smooth_quartz_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b SmoothQuartzStairs) New(props BlockProperties) Block {
	return SmoothQuartzStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}