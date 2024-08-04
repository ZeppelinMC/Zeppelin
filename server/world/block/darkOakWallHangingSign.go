package block

import (
	"strconv"
)

type DarkOakWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b DarkOakWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DarkOakWallHangingSign) New(props BlockProperties) Block {
	return DarkOakWallHangingSign{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}