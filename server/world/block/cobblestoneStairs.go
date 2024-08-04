package block

import (
	"strconv"
)

type CobblestoneStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b CobblestoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:cobblestone_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b CobblestoneStairs) New(props BlockProperties) Block {
	return CobblestoneStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}