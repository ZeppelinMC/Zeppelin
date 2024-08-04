package block

import (
	"strconv"
)

type PurpurStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b PurpurStairs) Encode() (string, BlockProperties) {
	return "minecraft:purpur_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b PurpurStairs) New(props BlockProperties) Block {
	return PurpurStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}