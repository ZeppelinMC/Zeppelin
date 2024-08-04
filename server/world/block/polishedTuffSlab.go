package block

import (
	"strconv"
)

type PolishedTuffSlab struct {
	Type string
	Waterlogged bool
}

func (b PolishedTuffSlab) Encode() (string, BlockProperties) {
	return "minecraft:polished_tuff_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PolishedTuffSlab) New(props BlockProperties) Block {
	return PolishedTuffSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}