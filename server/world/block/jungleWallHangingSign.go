package block

import (
	"strconv"
)

type JungleWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b JungleWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:jungle_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b JungleWallHangingSign) New(props BlockProperties) Block {
	return JungleWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}