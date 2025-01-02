package block

import (
	"strconv"
)

type BlackstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b BlackstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:blackstone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BlackstoneSlab) New(props BlockProperties) Block {
	return BlackstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}