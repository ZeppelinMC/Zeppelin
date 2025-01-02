package block

import (
	"strconv"
)

type NetherBrickFence struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b NetherBrickFence) Encode() (string, BlockProperties) {
	return "minecraft:nether_brick_fence", BlockProperties{
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
	}
}

func (b NetherBrickFence) New(props BlockProperties) Block {
	return NetherBrickFence{
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}