package block

import (
	"strconv"
)

type CherryHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b CherryHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:cherry_hanging_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b CherryHangingSign) New(props BlockProperties) Block {
	return CherryHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}