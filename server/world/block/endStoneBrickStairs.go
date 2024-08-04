package block

import (
	"strconv"
)

type EndStoneBrickStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b EndStoneBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:end_stone_brick_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b EndStoneBrickStairs) New(props BlockProperties) Block {
	return EndStoneBrickStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}