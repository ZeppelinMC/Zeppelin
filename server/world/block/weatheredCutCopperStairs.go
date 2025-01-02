package block

import (
	"strconv"
)

type WeatheredCutCopperStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b WeatheredCutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:weathered_cut_copper_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WeatheredCutCopperStairs) New(props BlockProperties) Block {
	return WeatheredCutCopperStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}