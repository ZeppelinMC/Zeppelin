package block

import (
	"strconv"
)

type QuartzStairs struct {
	Shape string
	Waterlogged bool
	Facing string
	Half string
}

func (b QuartzStairs) Encode() (string, BlockProperties) {
	return "minecraft:quartz_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b QuartzStairs) New(props BlockProperties) Block {
	return QuartzStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}