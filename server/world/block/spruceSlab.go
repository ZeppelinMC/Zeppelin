package block

import (
	"strconv"
)

type SpruceSlab struct {
	Waterlogged bool
	Type string
}

func (b SpruceSlab) Encode() (string, BlockProperties) {
	return "minecraft:spruce_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b SpruceSlab) New(props BlockProperties) Block {
	return SpruceSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}