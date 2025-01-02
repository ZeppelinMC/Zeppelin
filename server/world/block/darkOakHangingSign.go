package block

import (
	"strconv"
)

type DarkOakHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b DarkOakHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_hanging_sign", BlockProperties{
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DarkOakHangingSign) New(props BlockProperties) Block {
	return DarkOakHangingSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
		Attached: props["attached"] != "false",
	}
}