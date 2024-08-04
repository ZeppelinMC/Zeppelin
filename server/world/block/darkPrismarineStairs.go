package block

import (
	"strconv"
)

type DarkPrismarineStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b DarkPrismarineStairs) Encode() (string, BlockProperties) {
	return "minecraft:dark_prismarine_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b DarkPrismarineStairs) New(props BlockProperties) Block {
	return DarkPrismarineStairs{
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}