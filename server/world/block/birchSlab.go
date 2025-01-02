package block

import (
	"strconv"
)

type BirchSlab struct {
	Type string
	Waterlogged bool
}

func (b BirchSlab) Encode() (string, BlockProperties) {
	return "minecraft:birch_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BirchSlab) New(props BlockProperties) Block {
	return BirchSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}