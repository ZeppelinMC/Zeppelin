package block

import (
	"strconv"
)

type JungleSign struct {
	Waterlogged bool
	Rotation int
}

func (b JungleSign) Encode() (string, BlockProperties) {
	return "minecraft:jungle_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b JungleSign) New(props BlockProperties) Block {
	return JungleSign{
		Waterlogged: props["waterlogged"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}