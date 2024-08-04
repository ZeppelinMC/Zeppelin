package block

import (
	"strconv"
)

type JungleWallSign struct {
	Waterlogged bool
	Facing string
}

func (b JungleWallSign) Encode() (string, BlockProperties) {
	return "minecraft:jungle_wall_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b JungleWallSign) New(props BlockProperties) Block {
	return JungleWallSign{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}