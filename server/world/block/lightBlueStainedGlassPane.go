package block

import (
	"strconv"
)

type LightBlueStainedGlassPane struct {
	North bool
	South bool
	Waterlogged bool
	West bool
	East bool
}

func (b LightBlueStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_stained_glass_pane", BlockProperties{
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
	}
}

func (b LightBlueStainedGlassPane) New(props BlockProperties) Block {
	return LightBlueStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}