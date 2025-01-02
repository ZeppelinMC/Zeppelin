package block

import (
	"strconv"
)

type MossyCobblestoneStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b MossyCobblestoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:mossy_cobblestone_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b MossyCobblestoneStairs) New(props BlockProperties) Block {
	return MossyCobblestoneStairs{
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}