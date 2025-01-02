package block

import (
	"strconv"
)

type StoneBrickSlab struct {
	Waterlogged bool
	Type string
}

func (b StoneBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:stone_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b StoneBrickSlab) New(props BlockProperties) Block {
	return StoneBrickSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}