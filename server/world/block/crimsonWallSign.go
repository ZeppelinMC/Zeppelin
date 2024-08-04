package block

import (
	"strconv"
)

type CrimsonWallSign struct {
	Waterlogged bool
	Facing string
}

func (b CrimsonWallSign) Encode() (string, BlockProperties) {
	return "minecraft:crimson_wall_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b CrimsonWallSign) New(props BlockProperties) Block {
	return CrimsonWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}