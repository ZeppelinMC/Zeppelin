package block

import (
	"strconv"
)

type PolishedDeepslateStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b PolishedDeepslateStairs) Encode() (string, BlockProperties) {
	return "minecraft:polished_deepslate_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b PolishedDeepslateStairs) New(props BlockProperties) Block {
	return PolishedDeepslateStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}