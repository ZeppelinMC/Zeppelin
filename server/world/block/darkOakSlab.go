package block

import (
	"strconv"
)

type DarkOakSlab struct {
	Waterlogged bool
	Type string
}

func (b DarkOakSlab) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b DarkOakSlab) New(props BlockProperties) Block {
	return DarkOakSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}