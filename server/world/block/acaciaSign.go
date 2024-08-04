package block

import (
	"strconv"
)

type AcaciaSign struct {
	Rotation int
	Waterlogged bool
}

func (b AcaciaSign) Encode() (string, BlockProperties) {
	return "minecraft:acacia_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AcaciaSign) New(props BlockProperties) Block {
	return AcaciaSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}