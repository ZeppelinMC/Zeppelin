package block

import (
	"strconv"
)

type WaxedExposedCutCopperStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b WaxedExposedCutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_cut_copper_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b WaxedExposedCutCopperStairs) New(props BlockProperties) Block {
	return WaxedExposedCutCopperStairs{
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}