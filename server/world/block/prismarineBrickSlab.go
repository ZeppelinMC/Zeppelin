package block

import (
	"strconv"
)

type PrismarineBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b PrismarineBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:prismarine_brick_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b PrismarineBrickSlab) New(props BlockProperties) Block {
	return PrismarineBrickSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}