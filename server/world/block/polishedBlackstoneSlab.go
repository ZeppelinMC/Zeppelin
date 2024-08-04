package block

import (
	"strconv"
)

type PolishedBlackstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b PolishedBlackstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedBlackstoneSlab) New(props BlockProperties) Block {
	return PolishedBlackstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}