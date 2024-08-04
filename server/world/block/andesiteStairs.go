package block

import (
	"strconv"
)

type AndesiteStairs struct {
	Shape string
	Waterlogged bool
	Facing string
	Half string
}

func (b AndesiteStairs) Encode() (string, BlockProperties) {
	return "minecraft:andesite_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AndesiteStairs) New(props BlockProperties) Block {
	return AndesiteStairs{
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}