package block

import (
	"strconv"
)

type SpruceHangingSign struct {
	Attached bool
	Rotation int
	Waterlogged bool
}

func (b SpruceHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:spruce_hanging_sign", BlockProperties{
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SpruceHangingSign) New(props BlockProperties) Block {
	return SpruceHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}