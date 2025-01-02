package block

import (
	"strconv"
)

type NetherBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b NetherBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:nether_brick_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b NetherBrickSlab) New(props BlockProperties) Block {
	return NetherBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}