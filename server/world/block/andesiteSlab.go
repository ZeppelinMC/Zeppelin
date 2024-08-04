package block

import (
	"strconv"
)

type AndesiteSlab struct {
	Waterlogged bool
	Type string
}

func (b AndesiteSlab) Encode() (string, BlockProperties) {
	return "minecraft:andesite_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b AndesiteSlab) New(props BlockProperties) Block {
	return AndesiteSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}