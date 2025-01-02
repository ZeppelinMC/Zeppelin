package block

import (
	"strconv"
)

type EndStoneBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b EndStoneBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:end_stone_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b EndStoneBrickSlab) New(props BlockProperties) Block {
	return EndStoneBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}