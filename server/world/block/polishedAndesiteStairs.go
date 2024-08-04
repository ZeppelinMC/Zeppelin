package block

import (
	"strconv"
)

type PolishedAndesiteStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b PolishedAndesiteStairs) Encode() (string, BlockProperties) {
	return "minecraft:polished_andesite_stairs", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedAndesiteStairs) New(props BlockProperties) Block {
	return PolishedAndesiteStairs{
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}