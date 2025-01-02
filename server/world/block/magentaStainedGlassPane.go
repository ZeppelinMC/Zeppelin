package block

import (
	"strconv"
)

type MagentaStainedGlassPane struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b MagentaStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:magenta_stained_glass_pane", BlockProperties{
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MagentaStainedGlassPane) New(props BlockProperties) Block {
	return MagentaStainedGlassPane{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}