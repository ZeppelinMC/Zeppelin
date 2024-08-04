package block

import (
	"strconv"
)

type PrismarineSlab struct {
	Type string
	Waterlogged bool
}

func (b PrismarineSlab) Encode() (string, BlockProperties) {
	return "minecraft:prismarine_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b PrismarineSlab) New(props BlockProperties) Block {
	return PrismarineSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}