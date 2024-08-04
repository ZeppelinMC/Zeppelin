package block

import (
	"strconv"
)

type MangroveSlab struct {
	Type string
	Waterlogged bool
}

func (b MangroveSlab) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b MangroveSlab) New(props BlockProperties) Block {
	return MangroveSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}