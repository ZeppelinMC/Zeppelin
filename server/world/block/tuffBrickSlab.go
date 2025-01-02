package block

import (
	"strconv"
)

type TuffBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b TuffBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:tuff_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b TuffBrickSlab) New(props BlockProperties) Block {
	return TuffBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}