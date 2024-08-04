package block

import (
	"strconv"
)

type StoneStairs struct {
	Shape string
	Waterlogged bool
	Facing string
	Half string
}

func (b StoneStairs) Encode() (string, BlockProperties) {
	return "minecraft:stone_stairs", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b StoneStairs) New(props BlockProperties) Block {
	return StoneStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}