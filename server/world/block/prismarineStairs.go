package block

import (
	"strconv"
)

type PrismarineStairs struct {
	Shape string
	Waterlogged bool
	Facing string
	Half string
}

func (b PrismarineStairs) Encode() (string, BlockProperties) {
	return "minecraft:prismarine_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b PrismarineStairs) New(props BlockProperties) Block {
	return PrismarineStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}