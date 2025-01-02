package block

import (
	"strconv"
)

type DarkOakWallSign struct {
	Facing string
	Waterlogged bool
}

func (b DarkOakWallSign) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DarkOakWallSign) New(props BlockProperties) Block {
	return DarkOakWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}