package block

import (
	"strconv"
)

type CrimsonHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b CrimsonHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:crimson_hanging_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"attached": strconv.FormatBool(b.Attached),
	}
}

func (b CrimsonHangingSign) New(props BlockProperties) Block {
	return CrimsonHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}