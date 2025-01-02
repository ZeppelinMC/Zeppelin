package block

import (
	"strconv"
)

type StoneSlab struct {
	Type string
	Waterlogged bool
}

func (b StoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:stone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b StoneSlab) New(props BlockProperties) Block {
	return StoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}