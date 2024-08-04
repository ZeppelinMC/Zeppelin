package block

import (
	"strconv"
)

type BambooSlab struct {
	Type string
	Waterlogged bool
}

func (b BambooSlab) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BambooSlab) New(props BlockProperties) Block {
	return BambooSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}