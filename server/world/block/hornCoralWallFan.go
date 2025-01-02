package block

import (
	"strconv"
)

type HornCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b HornCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:horn_coral_wall_fan", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b HornCoralWallFan) New(props BlockProperties) Block {
	return HornCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}