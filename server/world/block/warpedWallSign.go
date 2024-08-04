package block

import (
	"strconv"
)

type WarpedWallSign struct {
	Facing string
	Waterlogged bool
}

func (b WarpedWallSign) Encode() (string, BlockProperties) {
	return "minecraft:warped_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WarpedWallSign) New(props BlockProperties) Block {
	return WarpedWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}