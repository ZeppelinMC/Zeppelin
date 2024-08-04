package block

import (
	"strconv"
)

type ExposedCutCopperStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b ExposedCutCopperStairs) Encode() (string, BlockProperties) {
	return "minecraft:exposed_cut_copper_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b ExposedCutCopperStairs) New(props BlockProperties) Block {
	return ExposedCutCopperStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}