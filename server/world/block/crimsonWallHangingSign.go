package block

import (
	"strconv"
)

type CrimsonWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b CrimsonWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:crimson_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CrimsonWallHangingSign) New(props BlockProperties) Block {
	return CrimsonWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}