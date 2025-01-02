package block

import (
	"strconv"
)

type HornCoralFan struct {
	Waterlogged bool
}

func (b HornCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:horn_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b HornCoralFan) New(props BlockProperties) Block {
	return HornCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}