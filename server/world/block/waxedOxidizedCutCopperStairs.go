package block

import (
	"strconv"
)

type WaxedOxidizedCutCopperStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b WaxedOxidizedCutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_cut_copper_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WaxedOxidizedCutCopperStairs) New(props BlockProperties) Block {
	return WaxedOxidizedCutCopperStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}