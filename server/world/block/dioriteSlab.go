package block

import (
	"strconv"
)

type DioriteSlab struct {
	Type string
	Waterlogged bool
}

func (b DioriteSlab) Encode() (string, BlockProperties) {
	return "minecraft:diorite_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DioriteSlab) New(props BlockProperties) Block {
	return DioriteSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}