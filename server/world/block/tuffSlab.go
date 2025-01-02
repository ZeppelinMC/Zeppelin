package block

import (
	"strconv"
)

type TuffSlab struct {
	Type string
	Waterlogged bool
}

func (b TuffSlab) Encode() (string, BlockProperties) {
	return "minecraft:tuff_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b TuffSlab) New(props BlockProperties) Block {
	return TuffSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}