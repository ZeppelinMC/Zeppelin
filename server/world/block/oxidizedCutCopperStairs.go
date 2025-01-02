package block

import (
	"strconv"
)

type OxidizedCutCopperStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b OxidizedCutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_cut_copper_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b OxidizedCutCopperStairs) New(props BlockProperties) Block {
	return OxidizedCutCopperStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}