package block

import (
	"strconv"
)

type PolishedAndesiteSlab struct {
	Type string
	Waterlogged bool
}

func (b PolishedAndesiteSlab) Encode() (string, BlockProperties) {
	return "minecraft:polished_andesite_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedAndesiteSlab) New(props BlockProperties) Block {
	return PolishedAndesiteSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}