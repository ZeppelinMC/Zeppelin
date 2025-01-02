package block

import (
	"strconv"
)

type BambooWallHangingSign struct {
	Facing string
	Waterlogged bool
}

func (b BambooWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BambooWallHangingSign) New(props BlockProperties) Block {
	return BambooWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}