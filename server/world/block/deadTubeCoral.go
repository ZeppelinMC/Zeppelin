package block

import (
	"strconv"
)

type DeadTubeCoral struct {
	Waterlogged bool
}

func (b DeadTubeCoral) Encode() (string, BlockProperties) {
	return "minecraft:dead_tube_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadTubeCoral) New(props BlockProperties) Block {
	return DeadTubeCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}