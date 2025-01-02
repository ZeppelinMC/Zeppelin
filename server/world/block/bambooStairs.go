package block

import (
	"strconv"
)

type BambooStairs struct {
	Waterlogged bool
	Facing string
	Half string
	Shape string
}

func (b BambooStairs) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_stairs", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"shape": b.Shape,
	}
}

func (b BambooStairs) New(props BlockProperties) Block {
	return BambooStairs{
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}