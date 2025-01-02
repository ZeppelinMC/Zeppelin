package block

import (
	"strconv"
)

type BambooFence struct {
	South bool
	Waterlogged bool
	West bool
	East bool
	North bool
}

func (b BambooFence) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_fence", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b BambooFence) New(props BlockProperties) Block {
	return BambooFence{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}