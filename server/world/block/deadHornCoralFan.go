package block

import (
	"strconv"
)

type DeadHornCoralFan struct {
	Waterlogged bool
}

func (b DeadHornCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_horn_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadHornCoralFan) New(props BlockProperties) Block {
	return DeadHornCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}