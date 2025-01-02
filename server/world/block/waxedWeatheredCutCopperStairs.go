package block

import (
	"strconv"
)

type WaxedWeatheredCutCopperStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b WaxedWeatheredCutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_cut_copper_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedWeatheredCutCopperStairs) New(props BlockProperties) Block {
	return WaxedWeatheredCutCopperStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}