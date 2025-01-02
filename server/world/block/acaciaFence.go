package block

import (
	"strconv"
)

type AcaciaFence struct {
	South bool
	Waterlogged bool
	West bool
	East bool
	North bool
}

func (b AcaciaFence) Encode() (string, BlockProperties) {
	return "minecraft:acacia_fence", BlockProperties{
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
	}
}

func (b AcaciaFence) New(props BlockProperties) Block {
	return AcaciaFence{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}