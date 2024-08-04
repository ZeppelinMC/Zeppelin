package block

import (
	"strconv"
)

type DeadFireCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b DeadFireCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_fire_coral_wall_fan", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadFireCoralWallFan) New(props BlockProperties) Block {
	return DeadFireCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}