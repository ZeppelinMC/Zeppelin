package block

import (
	"strconv"
)

type DarkOakFence struct {
	North bool
	South bool
	Waterlogged bool
	West bool
	East bool
}

func (b DarkOakFence) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_fence", BlockProperties{
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DarkOakFence) New(props BlockProperties) Block {
	return DarkOakFence{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}