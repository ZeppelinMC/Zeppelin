package block

import (
	"strconv"
)

type DeadFireCoral struct {
	Waterlogged bool
}

func (b DeadFireCoral) Encode() (string, BlockProperties) {
	return "minecraft:dead_fire_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeadFireCoral) New(props BlockProperties) Block {
	return DeadFireCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}