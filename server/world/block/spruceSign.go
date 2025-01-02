package block

import (
	"strconv"
)

type SpruceSign struct {
	Waterlogged bool
	Rotation int
}

func (b SpruceSign) Encode() (string, BlockProperties) {
	return "minecraft:spruce_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SpruceSign) New(props BlockProperties) Block {
	return SpruceSign{
		Waterlogged: props["waterlogged"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}