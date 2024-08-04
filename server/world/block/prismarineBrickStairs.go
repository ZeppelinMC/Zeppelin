package block

import (
	"strconv"
)

type PrismarineBrickStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b PrismarineBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:prismarine_brick_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PrismarineBrickStairs) New(props BlockProperties) Block {
	return PrismarineBrickStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}