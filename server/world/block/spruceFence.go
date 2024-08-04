package block

import (
	"strconv"
)

type SpruceFence struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b SpruceFence) Encode() (string, BlockProperties) {
	return "minecraft:spruce_fence", BlockProperties{
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SpruceFence) New(props BlockProperties) Block {
	return SpruceFence{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
	}
}