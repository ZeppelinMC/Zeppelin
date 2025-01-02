package block

import (
	"strconv"
)

type BrainCoralWallFan struct {
	Waterlogged bool
	Facing string
}

func (b BrainCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:brain_coral_wall_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b BrainCoralWallFan) New(props BlockProperties) Block {
	return BrainCoralWallFan{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}