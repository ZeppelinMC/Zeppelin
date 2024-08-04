package block

import (
	"strconv"
)

type OakWallHangingSign struct {
	Waterlogged bool
	Facing string
}

func (b OakWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:oak_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OakWallHangingSign) New(props BlockProperties) Block {
	return OakWallHangingSign{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}