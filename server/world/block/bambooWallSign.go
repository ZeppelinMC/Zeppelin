package block

import (
	"strconv"
)

type BambooWallSign struct {
	Facing string
	Waterlogged bool
}

func (b BambooWallSign) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BambooWallSign) New(props BlockProperties) Block {
	return BambooWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}