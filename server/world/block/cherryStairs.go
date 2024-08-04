package block

import (
	"strconv"
)

type CherryStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b CherryStairs) Encode() (string, BlockProperties) {
	return "minecraft:cherry_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b CherryStairs) New(props BlockProperties) Block {
	return CherryStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}