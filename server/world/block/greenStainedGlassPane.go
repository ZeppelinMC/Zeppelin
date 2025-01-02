package block

import (
	"strconv"
)

type GreenStainedGlassPane struct {
	South bool
	Waterlogged bool
	West bool
	East bool
	North bool
}

func (b GreenStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:green_stained_glass_pane", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
	}
}

func (b GreenStainedGlassPane) New(props BlockProperties) Block {
	return GreenStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}