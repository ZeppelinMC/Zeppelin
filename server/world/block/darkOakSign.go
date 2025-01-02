package block

import (
	"strconv"
)

type DarkOakSign struct {
	Waterlogged bool
	Rotation int
}

func (b DarkOakSign) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b DarkOakSign) New(props BlockProperties) Block {
	return DarkOakSign{
		Waterlogged: props["waterlogged"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}