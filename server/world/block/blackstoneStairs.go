package block

import (
	"strconv"
)

type BlackstoneStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b BlackstoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:blackstone_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BlackstoneStairs) New(props BlockProperties) Block {
	return BlackstoneStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}