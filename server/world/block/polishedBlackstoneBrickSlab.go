package block

import (
	"strconv"
)

type PolishedBlackstoneBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b PolishedBlackstoneBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedBlackstoneBrickSlab) New(props BlockProperties) Block {
	return PolishedBlackstoneBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}