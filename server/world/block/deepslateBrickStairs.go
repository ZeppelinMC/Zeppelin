package block

import (
	"strconv"
)

type DeepslateBrickStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b DeepslateBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_brick_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeepslateBrickStairs) New(props BlockProperties) Block {
	return DeepslateBrickStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}