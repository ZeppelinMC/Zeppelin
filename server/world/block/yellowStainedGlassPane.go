package block

import (
	"strconv"
)

type YellowStainedGlassPane struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b YellowStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:yellow_stained_glass_pane", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
	}
}

func (b YellowStainedGlassPane) New(props BlockProperties) Block {
	return YellowStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}