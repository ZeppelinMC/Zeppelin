package block

import (
	"strconv"
)

type DeadBubbleCoralWallFan struct {
	Facing string
	Waterlogged bool
}

func (b DeadBubbleCoralWallFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_bubble_coral_wall_fan", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadBubbleCoralWallFan) New(props BlockProperties) Block {
	return DeadBubbleCoralWallFan{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}