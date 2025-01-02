package block

import (
	"strconv"
)

type DeadBubbleCoralFan struct {
	Waterlogged bool
}

func (b DeadBubbleCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_bubble_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadBubbleCoralFan) New(props BlockProperties) Block {
	return DeadBubbleCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}