package block

import (
	"strconv"
)

type PolishedGraniteSlab struct {
	Type string
	Waterlogged bool
}

func (b PolishedGraniteSlab) Encode() (string, BlockProperties) {
	return "minecraft:polished_granite_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedGraniteSlab) New(props BlockProperties) Block {
	return PolishedGraniteSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}