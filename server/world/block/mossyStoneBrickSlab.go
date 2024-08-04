package block

import (
	"strconv"
)

type MossyStoneBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b MossyStoneBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:mossy_stone_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MossyStoneBrickSlab) New(props BlockProperties) Block {
	return MossyStoneBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}