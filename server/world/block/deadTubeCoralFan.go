package block

import (
	"strconv"
)

type DeadTubeCoralFan struct {
	Waterlogged bool
}

func (b DeadTubeCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_tube_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadTubeCoralFan) New(props BlockProperties) Block {
	return DeadTubeCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}