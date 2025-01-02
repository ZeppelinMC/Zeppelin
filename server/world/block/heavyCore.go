package block

import (
	"strconv"
)

type HeavyCore struct {
	Waterlogged bool
}

func (b HeavyCore) Encode() (string, BlockProperties) {
	return "minecraft:heavy_core", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b HeavyCore) New(props BlockProperties) Block {
	return HeavyCore{
		Waterlogged: props["waterlogged"] != "false",
	}
}