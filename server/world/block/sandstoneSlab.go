package block

import (
	"strconv"
)

type SandstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b SandstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:sandstone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SandstoneSlab) New(props BlockProperties) Block {
	return SandstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}