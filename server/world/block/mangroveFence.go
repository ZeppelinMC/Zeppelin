package block

import (
	"strconv"
)

type MangroveFence struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b MangroveFence) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_fence", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b MangroveFence) New(props BlockProperties) Block {
	return MangroveFence{
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
	}
}