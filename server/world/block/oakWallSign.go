package block

import (
	"strconv"
)

type OakWallSign struct {
	Waterlogged bool
	Facing string
}

func (b OakWallSign) Encode() (string, BlockProperties) {
	return "minecraft:oak_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OakWallSign) New(props BlockProperties) Block {
	return OakWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}