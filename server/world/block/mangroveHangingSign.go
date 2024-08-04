package block

import (
	"strconv"
)

type MangroveHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b MangroveHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_hanging_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"attached": strconv.FormatBool(b.Attached),
	}
}

func (b MangroveHangingSign) New(props BlockProperties) Block {
	return MangroveHangingSign{
		Waterlogged: props["waterlogged"] != "false",
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}