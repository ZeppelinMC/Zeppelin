package block

import (
	"strconv"
)

type BirchSign struct {
	Rotation int
	Waterlogged bool
}

func (b BirchSign) Encode() (string, BlockProperties) {
	return "minecraft:birch_sign", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b BirchSign) New(props BlockProperties) Block {
	return BirchSign{
		Waterlogged: props["waterlogged"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}