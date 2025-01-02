package block

import (
	"strconv"
)

type CrimsonSign struct {
	Rotation int
	Waterlogged bool
}

func (b CrimsonSign) Encode() (string, BlockProperties) {
	return "minecraft:crimson_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CrimsonSign) New(props BlockProperties) Block {
	return CrimsonSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}