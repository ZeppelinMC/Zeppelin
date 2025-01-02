package block

import (
	"strconv"
)

type BirchWallSign struct {
	Facing string
	Waterlogged bool
}

func (b BirchWallSign) Encode() (string, BlockProperties) {
	return "minecraft:birch_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BirchWallSign) New(props BlockProperties) Block {
	return BirchWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}