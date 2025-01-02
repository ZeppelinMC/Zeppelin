package block

import (
	"strconv"
)

type MangroveSign struct {
	Rotation int
	Waterlogged bool
}

func (b MangroveSign) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MangroveSign) New(props BlockProperties) Block {
	return MangroveSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}