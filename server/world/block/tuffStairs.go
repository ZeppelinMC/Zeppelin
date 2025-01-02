package block

import (
	"strconv"
)

type TuffStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b TuffStairs) Encode() (string, BlockProperties) {
	return "minecraft:tuff_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b TuffStairs) New(props BlockProperties) Block {
	return TuffStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}