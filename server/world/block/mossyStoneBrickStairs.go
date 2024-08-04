package block

import (
	"strconv"
)

type MossyStoneBrickStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b MossyStoneBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:mossy_stone_brick_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MossyStoneBrickStairs) New(props BlockProperties) Block {
	return MossyStoneBrickStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}