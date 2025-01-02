package block

import (
	"strconv"
)

type DeadHornCoral struct {
	Waterlogged bool
}

func (b DeadHornCoral) Encode() (string, BlockProperties) {
	return "minecraft:dead_horn_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadHornCoral) New(props BlockProperties) Block {
	return DeadHornCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}