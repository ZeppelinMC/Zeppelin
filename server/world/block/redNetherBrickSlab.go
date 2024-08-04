package block

import (
	"strconv"
)

type RedNetherBrickSlab struct {
	Waterlogged bool
	Type string
}

func (b RedNetherBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:red_nether_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b RedNetherBrickSlab) New(props BlockProperties) Block {
	return RedNetherBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}