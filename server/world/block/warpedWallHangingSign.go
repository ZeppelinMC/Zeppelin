package block

import (
	"strconv"
)

type WarpedWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b WarpedWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:warped_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WarpedWallHangingSign) New(props BlockProperties) Block {
	return WarpedWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}