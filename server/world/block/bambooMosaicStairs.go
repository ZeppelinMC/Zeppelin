package block

import (
	"strconv"
)

type BambooMosaicStairs struct {
	Half string
	Shape string
	Waterlogged bool
	Facing string
}

func (b BambooMosaicStairs) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_mosaic_stairs", BlockProperties{
		"half": b.Half,
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b BambooMosaicStairs) New(props BlockProperties) Block {
	return BambooMosaicStairs{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}