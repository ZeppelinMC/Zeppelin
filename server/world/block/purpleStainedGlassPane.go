package block

import (
	"strconv"
)

type PurpleStainedGlassPane struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b PurpleStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:purple_stained_glass_pane", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
	}
}

func (b PurpleStainedGlassPane) New(props BlockProperties) Block {
	return PurpleStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}