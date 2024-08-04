package block

import (
	"strconv"
)

type DeadBrainCoralFan struct {
	Waterlogged bool
}

func (b DeadBrainCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_brain_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadBrainCoralFan) New(props BlockProperties) Block {
	return DeadBrainCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}