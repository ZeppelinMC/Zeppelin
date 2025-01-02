package block

import (
	"strconv"
)

type TubeCoralFan struct {
	Waterlogged bool
}

func (b TubeCoralFan) Encode() (string, BlockProperties) {
	return "minecraft:tube_coral_fan", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b TubeCoralFan) New(props BlockProperties) Block {
	return TubeCoralFan{
		Waterlogged: props["waterlogged"] != "false",
	}
}