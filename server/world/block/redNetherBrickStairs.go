package block

import (
	"strconv"
)

type RedNetherBrickStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b RedNetherBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:red_nether_brick_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b RedNetherBrickStairs) New(props BlockProperties) Block {
	return RedNetherBrickStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}