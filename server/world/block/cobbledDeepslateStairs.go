package block

import (
	"strconv"
)

type CobbledDeepslateStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b CobbledDeepslateStairs) Encode() (string, BlockProperties) {
	return "minecraft:cobbled_deepslate_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CobbledDeepslateStairs) New(props BlockProperties) Block {
	return CobbledDeepslateStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}