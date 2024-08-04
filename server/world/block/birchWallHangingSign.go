package block

import (
	"strconv"
)

type BirchWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b BirchWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:birch_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BirchWallHangingSign) New(props BlockProperties) Block {
	return BirchWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}