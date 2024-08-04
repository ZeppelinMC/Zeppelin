package block

import (
	"strconv"
)

type PolishedGraniteStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b PolishedGraniteStairs) Encode() (string, BlockProperties) {
	return "minecraft:polished_granite_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b PolishedGraniteStairs) New(props BlockProperties) Block {
	return PolishedGraniteStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}