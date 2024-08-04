package block

import (
	"strconv"
)

type WarpedFence struct {
	West bool
	East bool
	North bool
	South bool
	Waterlogged bool
}

func (b WarpedFence) Encode() (string, BlockProperties) {
	return "minecraft:warped_fence", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
	}
}

func (b WarpedFence) New(props BlockProperties) Block {
	return WarpedFence{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}