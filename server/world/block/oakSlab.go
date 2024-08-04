package block

import (
	"strconv"
)

type OakSlab struct {
	Type string
	Waterlogged bool
}

func (b OakSlab) Encode() (string, BlockProperties) {
	return "minecraft:oak_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OakSlab) New(props BlockProperties) Block {
	return OakSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}