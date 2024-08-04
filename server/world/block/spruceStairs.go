package block

import (
	"strconv"
)

type SpruceStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b SpruceStairs) Encode() (string, BlockProperties) {
	return "minecraft:spruce_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b SpruceStairs) New(props BlockProperties) Block {
	return SpruceStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}