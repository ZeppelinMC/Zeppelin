package block

import (
	"strconv"
)

type OakFence struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b OakFence) Encode() (string, BlockProperties) {
	return "minecraft:oak_fence", BlockProperties{
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
	}
}

func (b OakFence) New(props BlockProperties) Block {
	return OakFence{
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
	}
}