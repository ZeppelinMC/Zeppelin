package block

import (
	"strconv"
)

type DarkOakStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b DarkOakStairs) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DarkOakStairs) New(props BlockProperties) Block {
	return DarkOakStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}