package block

import (
	"strconv"
)

type CyanStainedGlassPane struct {
	Waterlogged bool
	West bool
	East bool
	North bool
	South bool
}

func (b CyanStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:cyan_stained_glass_pane", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b CyanStainedGlassPane) New(props BlockProperties) Block {
	return CyanStainedGlassPane{
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
	}
}