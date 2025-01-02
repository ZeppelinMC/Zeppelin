package block

import (
	"strconv"
)

type BambooSign struct {
	Waterlogged bool
	Rotation int
}

func (b BambooSign) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_sign", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BambooSign) New(props BlockProperties) Block {
	return BambooSign{
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}