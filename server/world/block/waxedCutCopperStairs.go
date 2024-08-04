package block

import (
	"strconv"
)

type WaxedCutCopperStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b WaxedCutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:waxed_cut_copper_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedCutCopperStairs) New(props BlockProperties) Block {
	return WaxedCutCopperStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}