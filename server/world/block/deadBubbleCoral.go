package block

import (
	"strconv"
)

type DeadBubbleCoral struct {
	Waterlogged bool
}

func (b DeadBubbleCoral) Encode() (string, BlockProperties) {
	return "minecraft:dead_bubble_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadBubbleCoral) New(props BlockProperties) Block {
	return DeadBubbleCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}