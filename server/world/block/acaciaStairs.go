package block

import (
	"strconv"
)

type AcaciaStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b AcaciaStairs) Encode() (string, BlockProperties) {
	return "minecraft:acacia_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AcaciaStairs) New(props BlockProperties) Block {
	return AcaciaStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}