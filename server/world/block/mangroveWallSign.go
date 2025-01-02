package block

import (
	"strconv"
)

type MangroveWallSign struct {
	Facing string
	Waterlogged bool
}

func (b MangroveWallSign) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MangroveWallSign) New(props BlockProperties) Block {
	return MangroveWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}