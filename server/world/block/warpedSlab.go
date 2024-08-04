package block

import (
	"strconv"
)

type WarpedSlab struct {
	Type string
	Waterlogged bool
}

func (b WarpedSlab) Encode() (string, BlockProperties) {
	return "minecraft:warped_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WarpedSlab) New(props BlockProperties) Block {
	return WarpedSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}