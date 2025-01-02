package block

import (
	"strconv"
)

type DioriteStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b DioriteStairs) Encode() (string, BlockProperties) {
	return "minecraft:diorite_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DioriteStairs) New(props BlockProperties) Block {
	return DioriteStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}