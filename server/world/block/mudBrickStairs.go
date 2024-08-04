package block

import (
	"strconv"
)

type MudBrickStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b MudBrickStairs) Encode() (string, BlockProperties) {
	return "minecraft:mud_brick_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b MudBrickStairs) New(props BlockProperties) Block {
	return MudBrickStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}