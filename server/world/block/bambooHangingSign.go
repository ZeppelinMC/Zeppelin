package block

import (
	"strconv"
)

type BambooHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b BambooHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_hanging_sign", BlockProperties{
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BambooHangingSign) New(props BlockProperties) Block {
	return BambooHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}