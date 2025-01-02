package block

import (
	"strconv"
)

type JungleSlab struct {
	Type string
	Waterlogged bool
}

func (b JungleSlab) Encode() (string, BlockProperties) {
	return "minecraft:jungle_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b JungleSlab) New(props BlockProperties) Block {
	return JungleSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}