package block

import (
	"strconv"
)

type DeadHornCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b DeadHornCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_horn_coral_wall_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b DeadHornCoralWallFan) New(props BlockProperties) Block {
	return DeadHornCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}