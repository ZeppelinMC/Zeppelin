package block

import (
	"strconv"
)

type BambooMosaicSlab struct {
	Type string
	Waterlogged bool
}

func (b BambooMosaicSlab) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_mosaic_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BambooMosaicSlab) New(props BlockProperties) Block {
	return BambooMosaicSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}