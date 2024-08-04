package block

import (
	"strconv"
)

type GraniteSlab struct {
	Type string
	Waterlogged bool
}

func (b GraniteSlab) Encode() (string, BlockProperties) {
	return "minecraft:granite_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b GraniteSlab) New(props BlockProperties) Block {
	return GraniteSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}