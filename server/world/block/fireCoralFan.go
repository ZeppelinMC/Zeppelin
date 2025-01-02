package block

import (
	"strconv"
)

type FireCoralFan struct {
	Waterlogged bool
}

func (b FireCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:fire_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b FireCoralFan) New(props BlockProperties) Block {
	return FireCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}