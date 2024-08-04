package block

import (
	"strconv"
)

type SandstoneStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b SandstoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:sandstone_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SandstoneStairs) New(props BlockProperties) Block {
	return SandstoneStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}