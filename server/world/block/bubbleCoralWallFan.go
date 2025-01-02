package block

import (
	"strconv"
)

type BubbleCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b BubbleCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:bubble_coral_wall_fan", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BubbleCoralWallFan) New(props BlockProperties) Block {
	return BubbleCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}