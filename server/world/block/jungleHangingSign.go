package block

import (
	"strconv"
)

type JungleHangingSign struct {
	Rotation int
	Waterlogged bool
	Attached bool
}

func (b JungleHangingSign) Encode() (string, BlockProperties) {
	return "minecraft:jungle_hanging_sign", BlockProperties{
		"attached": strconv.FormatBool(b.Attached),
		"rotation": strconv.Itoa(b.Rotation),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b JungleHangingSign) New(props BlockProperties) Block {
	return JungleHangingSign{
		Attached: props["attached"] != "false",
		Rotation: atoi(props["rotation"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}