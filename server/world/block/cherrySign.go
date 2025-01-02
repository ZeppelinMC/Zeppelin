package block

import (
	"strconv"
)

type CherrySign struct {
	Rotation int
	Waterlogged bool
}

func (b CherrySign) Encode() (string, BlockProperties) {
	return "minecraft:cherry_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CherrySign) New(props BlockProperties) Block {
	return CherrySign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}