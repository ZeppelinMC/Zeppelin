package block

import (
	"strconv"
)

type CrimsonStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b CrimsonStairs) Encode() (string, BlockProperties) {
	return "minecraft:crimson_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b CrimsonStairs) New(props BlockProperties) Block {
	return CrimsonStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}