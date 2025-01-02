package block

import (
	"strconv"
)

type WarpedHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b WarpedHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:warped_hanging_sign", BlockProperties{
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WarpedHangingSign) New(props BlockProperties) Block {
	return WarpedHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}