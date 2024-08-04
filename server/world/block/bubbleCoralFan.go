package block

import (
	"strconv"
)

type BubbleCoralFan struct {
	Waterlogged bool
}

func (b BubbleCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:bubble_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BubbleCoralFan) New(props BlockProperties) Block {
	return BubbleCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}