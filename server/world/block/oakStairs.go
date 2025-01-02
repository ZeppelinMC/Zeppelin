package block

import (
	"strconv"
)

type OakStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b OakStairs) Encode() (string, BlockProperties) {
	return "minecraft:oak_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OakStairs) New(props BlockProperties) Block {
	return OakStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}