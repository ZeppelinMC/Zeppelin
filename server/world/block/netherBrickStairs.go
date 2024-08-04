package block

import (
	"strconv"
)

type NetherBrickStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b NetherBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:nether_brick_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b NetherBrickStairs) New(props BlockProperties) Block {
	return NetherBrickStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}