package block

import (
	"strconv"
)

type OakSign struct {
	Rotation int
	Waterlogged bool
}

func (b OakSign) Encode() (string, BlockProperties) {
	return "minecraft:oak_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OakSign) New(props BlockProperties) Block {
	return OakSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}