package block

import (
	"strconv"
)

type BlackStainedGlassPane struct {
	Waterlogged bool
	West bool
	East bool
	North bool
	South bool
}

func (b BlackStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:black_stained_glass_pane", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b BlackStainedGlassPane) New(props BlockProperties) Block {
	return BlackStainedGlassPane{
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
	}
}