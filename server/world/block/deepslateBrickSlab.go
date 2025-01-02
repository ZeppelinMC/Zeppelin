package block

import (
	"strconv"
)

type DeepslateBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b DeepslateBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeepslateBrickSlab) New(props BlockProperties) Block {
	return DeepslateBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}