package block

import (
	"strconv"
)

type AcaciaWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b AcaciaWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:acacia_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AcaciaWallHangingSign) New(props BlockProperties) Block {
	return AcaciaWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}