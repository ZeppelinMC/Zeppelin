package block

import (
	"strconv"
)

type DeadTubeCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b DeadTubeCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_tube_coral_wall_fan", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadTubeCoralWallFan) New(props BlockProperties) Block {
	return DeadTubeCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}