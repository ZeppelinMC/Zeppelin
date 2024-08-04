package block

import (
	"strconv"
)

type MangroveWallHangingSign struct {
	Waterlogged bool
	Facing string
}

func (b MangroveWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_wall_hanging_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b MangroveWallHangingSign) New(props BlockProperties) Block {
	return MangroveWallHangingSign{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}