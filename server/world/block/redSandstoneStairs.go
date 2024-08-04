package block

import (
	"strconv"
)

type RedSandstoneStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b RedSandstoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:red_sandstone_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b RedSandstoneStairs) New(props BlockProperties) Block {
	return RedSandstoneStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}