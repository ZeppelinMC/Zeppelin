package block

import (
	"strconv"
)

type TubeCoralWallFan struct {
	Waterlogged bool
	Facing string
}

func (b TubeCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:tube_coral_wall_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b TubeCoralWallFan) New(props BlockProperties) Block {
	return TubeCoralWallFan{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}