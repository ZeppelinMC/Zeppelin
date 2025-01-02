package block

import (
	"strconv"
)

type BrickSlab struct {
	Type string
	Waterlogged bool
}

func (b BrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BrickSlab) New(props BlockProperties) Block {
	return BrickSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}