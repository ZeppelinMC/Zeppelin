package block

import (
	"strconv"
)

type DeepslateTileSlab struct {
	Type string
	Waterlogged bool
}

func (b DeepslateTileSlab) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_tile_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeepslateTileSlab) New(props BlockProperties) Block {
	return DeepslateTileSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}