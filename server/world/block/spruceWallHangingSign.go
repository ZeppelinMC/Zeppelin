package block

import (
	"strconv"
)

type SpruceWallHangingSign struct {
	Waterlogged bool
	Facing string
}

func (b SpruceWallHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:spruce_wall_hanging_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SpruceWallHangingSign) New(props BlockProperties) Block {
	return SpruceWallHangingSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}