package block

import (
	"strconv"
)

type MangroveStairs struct {
	Facing string
	Half string
	Shape string
	Waterlogged bool
}

func (b MangroveStairs) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b MangroveStairs) New(props BlockProperties) Block {
	return MangroveStairs{
		Facing: props["facing"],
		Half: props["half"],
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}