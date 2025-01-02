package block

import (
	"strconv"
)

type WarpedSign struct {
	Waterlogged bool
	Rotation int
}

func (b WarpedSign) Encode() (string, BlockProperties) {
	return "minecraft:warped_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b WarpedSign) New(props BlockProperties) Block {
	return WarpedSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}