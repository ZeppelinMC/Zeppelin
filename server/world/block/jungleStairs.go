package block

import (
	"strconv"
)

type JungleStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b JungleStairs) Encode() (string, BlockProperties) {
	return "minecraft:jungle_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b JungleStairs) New(props BlockProperties) Block {
	return JungleStairs{
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}