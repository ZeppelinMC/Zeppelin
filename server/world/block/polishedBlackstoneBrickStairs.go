package block

import (
	"strconv"
)

type PolishedBlackstoneBrickStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b PolishedBlackstoneBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_brick_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b PolishedBlackstoneBrickStairs) New(props BlockProperties) Block {
	return PolishedBlackstoneBrickStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}