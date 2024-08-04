package block

import (
	"strconv"
)

type AcaciaHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b AcaciaHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:acacia_hanging_sign", BlockProperties{
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AcaciaHangingSign) New(props BlockProperties) Block {
	return AcaciaHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}