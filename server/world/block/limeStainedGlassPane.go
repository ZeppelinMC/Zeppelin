package block

import (
	"strconv"
)

type LimeStainedGlassPane struct {
	West bool
	East bool
	North bool
	South bool
	Waterlogged bool
}

func (b LimeStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:lime_stained_glass_pane", BlockProperties{
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
	}
}

func (b LimeStainedGlassPane) New(props BlockProperties) Block {
	return LimeStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}