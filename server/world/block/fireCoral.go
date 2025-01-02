package block

import (
	"strconv"
)

type FireCoral struct {
	Waterlogged bool
}

func (b FireCoral) Encode() (string, BlockProperties) {
	return "minecraft:fire_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b FireCoral) New(props BlockProperties) Block {
	return FireCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}