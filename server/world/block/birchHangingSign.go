package block

import (
	"strconv"
)

type BirchHangingSign struct {
	Waterlogged bool
	Attached bool
	Rotation int
}

func (b BirchHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:birch_hanging_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b BirchHangingSign) New(props BlockProperties) Block {
	return BirchHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}