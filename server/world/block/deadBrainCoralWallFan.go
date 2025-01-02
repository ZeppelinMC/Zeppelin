package block

import (
	"strconv"
)

type DeadBrainCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b DeadBrainCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_brain_coral_wall_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b DeadBrainCoralWallFan) New(props BlockProperties) Block {
	return DeadBrainCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}