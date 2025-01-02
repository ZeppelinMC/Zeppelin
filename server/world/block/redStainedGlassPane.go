package block

import (
	"strconv"
)

type RedStainedGlassPane struct {
	West bool
	East bool
	North bool
	South bool
	Waterlogged bool
}

func (b RedStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:red_stained_glass_pane", BlockProperties{
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b RedStainedGlassPane) New(props BlockProperties) Block {
	return RedStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}