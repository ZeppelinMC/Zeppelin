package block

import (
	"strconv"
)

type BrainCoralFan struct {
	Waterlogged bool
}

func (b BrainCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:brain_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BrainCoralFan) New(props BlockProperties) Block {
	return BrainCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}