package block

import (
	"strconv"
)

type BlueStainedGlassPane struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b BlueStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:blue_stained_glass_pane", BlockProperties{
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
	}
}

func (b BlueStainedGlassPane) New(props BlockProperties) Block {
	return BlueStainedGlassPane{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}