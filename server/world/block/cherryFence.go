package block

import (
	"strconv"
)

type CherryFence struct {
	Waterlogged bool
	West bool
	East bool
	North bool
	South bool
}

func (b CherryFence) Encode() (string, BlockProperties) {
	return "minecraft:cherry_fence", BlockProperties{
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
	}
}

func (b CherryFence) New(props BlockProperties) Block {
	return CherryFence{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
	}
}