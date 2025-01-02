package block

import (
	"strconv"
)

type AcaciaWallSign struct {
	Facing string
	Waterlogged bool
}

func (b AcaciaWallSign) Encode() (string, BlockProperties) {
	return "minecraft:acacia_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AcaciaWallSign) New(props BlockProperties) Block {
	return AcaciaWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}