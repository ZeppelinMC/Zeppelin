package block

import (
	"strconv"
)

type PolishedDioriteStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b PolishedDioriteStairs) Encode() (string, BlockProperties) {
	return "minecraft:polished_diorite_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedDioriteStairs) New(props BlockProperties) Block {
	return PolishedDioriteStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}