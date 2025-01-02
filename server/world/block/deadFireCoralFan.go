package block

import (
	"strconv"
)

type DeadFireCoralFan struct {
	Waterlogged bool
}

func (b DeadFireCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:dead_fire_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadFireCoralFan) New(props BlockProperties) Block {
	return DeadFireCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}