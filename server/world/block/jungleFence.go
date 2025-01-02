package block

import (
	"strconv"
)

type JungleFence struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b JungleFence) Encode() (string, BlockProperties) {
	return "minecraft:jungle_fence", BlockProperties{
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
	}
}

func (b JungleFence) New(props BlockProperties) Block {
	return JungleFence{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
	}
}