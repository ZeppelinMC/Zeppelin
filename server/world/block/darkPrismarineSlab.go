package block

import (
	"strconv"
)

type DarkPrismarineSlab struct {
	Type string
	Waterlogged bool
}

func (b DarkPrismarineSlab) Encode() (string, BlockProperties) {
	return "minecraft:dark_prismarine_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DarkPrismarineSlab) New(props BlockProperties) Block {
	return DarkPrismarineSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}