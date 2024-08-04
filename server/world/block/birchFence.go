package block

import (
	"strconv"
)

type BirchFence struct {
	West bool
	East bool
	North bool
	South bool
	Waterlogged bool
}

func (b BirchFence) Encode() (string, BlockProperties) {
	return "minecraft:birch_fence", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
	}
}

func (b BirchFence) New(props BlockProperties) Block {
	return BirchFence{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
	}
}