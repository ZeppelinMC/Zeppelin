package block

import (
	"strconv"
)

type CutCopperStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b CutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:cut_copper_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CutCopperStairs) New(props BlockProperties) Block {
	return CutCopperStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}