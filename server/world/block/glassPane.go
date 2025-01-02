package block

import (
	"strconv"
)

type GlassPane struct {
	Waterlogged bool
	West bool
	East bool
	North bool
	South bool
}

func (b GlassPane) Encode() (string, BlockProperties) {
	return "minecraft:glass_pane", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b GlassPane) New(props BlockProperties) Block {
	return GlassPane{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}