package block

import (
	"strconv"
)

type DeadBrainCoral struct {
	Waterlogged bool
}

func (b DeadBrainCoral) Encode() (string, BlockProperties) {
	return "minecraft:dead_brain_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadBrainCoral) New(props BlockProperties) Block {
	return DeadBrainCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}