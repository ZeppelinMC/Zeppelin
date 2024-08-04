package block

import (
	"strconv"
)

type FireCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b FireCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:fire_coral_wall_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b FireCoralWallFan) New(props BlockProperties) Block {
	return FireCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}