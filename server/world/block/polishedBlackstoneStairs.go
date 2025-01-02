package block

import (
	"strconv"
)

type PolishedBlackstoneStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b PolishedBlackstoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b PolishedBlackstoneStairs) New(props BlockProperties) Block {
	return PolishedBlackstoneStairs{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
	}
}