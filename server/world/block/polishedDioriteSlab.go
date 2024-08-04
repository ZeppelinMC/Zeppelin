package block

import (
	"strconv"
)

type PolishedDioriteSlab struct {
	Type string
	Waterlogged bool
}

func (b PolishedDioriteSlab) Encode() (string, BlockProperties) {
	return "minecraft:polished_diorite_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedDioriteSlab) New(props BlockProperties) Block {
	return PolishedDioriteSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}