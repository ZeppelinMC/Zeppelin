package block

import (
	"strconv"
)

type GraniteStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b GraniteStairs) Encode() (string, BlockProperties) {
	return "minecraft:granite_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b GraniteStairs) New(props BlockProperties) Block {
	return GraniteStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}