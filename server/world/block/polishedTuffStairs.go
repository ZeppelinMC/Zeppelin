package block

import (
	"strconv"
)

type PolishedTuffStairs struct {
	Shape string
	Waterlogged bool
	Facing string
	Half string
}

func (b PolishedTuffStairs) Encode() (string, BlockProperties) {
	return "minecraft:polished_tuff_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedTuffStairs) New(props BlockProperties) Block {
	return PolishedTuffStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}