package block

import (
	"strconv"
)

type SpruceWallSign struct {
	Facing string
	Waterlogged bool
}

func (b SpruceWallSign) Encode() (string, BlockProperties) {
	return "minecraft:spruce_wall_sign", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SpruceWallSign) New(props BlockProperties) Block {
	return SpruceWallSign{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}