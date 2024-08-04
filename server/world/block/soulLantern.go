package block

import (
	"strconv"
)

type SoulLantern struct {
	Hanging bool
	Waterlogged bool
}

func (b SoulLantern) Encode() (string, BlockProperties) {
	return "minecraft:soul_lantern", BlockProperties{
		"hanging": strconv.FormatBool(b.Hanging),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SoulLantern) New(props BlockProperties) Block {
	return SoulLantern{
		Hanging: props["hanging"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}