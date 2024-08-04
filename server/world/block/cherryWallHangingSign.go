package block

import (
	"strconv"
)

type CherryWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b CherryWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:cherry_wall_hanging_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b CherryWallHangingSign) New(props BlockProperties) Block {
	return CherryWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}