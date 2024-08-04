package block

import (
	"strconv"
)

type StoneBrickStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b StoneBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:stone_brick_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b StoneBrickStairs) New(props BlockProperties) Block {
	return StoneBrickStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}