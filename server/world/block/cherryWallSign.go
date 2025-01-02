package block

import (
	"strconv"
)

type CherryWallSign struct {
	Facing string
	Waterlogged bool
}

func (b CherryWallSign) Encode() (string, BlockProperties) {
	return "minecraft:cherry_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CherryWallSign) New(props BlockProperties) Block {
	return CherryWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}