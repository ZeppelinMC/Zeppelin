package block

import (
	"strconv"
)

type OakHangingSign struct {
	Waterlogged bool
	Attached bool
	Rotation int
}

func (b OakHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:oak_hanging_sign", BlockProperties{
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OakHangingSign) New(props BlockProperties) Block {
	return OakHangingSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
		Attached: props["attached"] != "false",
	}
}