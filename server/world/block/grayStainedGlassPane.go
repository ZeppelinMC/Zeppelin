package block

import (
	"strconv"
)

type GrayStainedGlassPane struct {
	West bool
	East bool
	North bool
	South bool
	Waterlogged bool
}

func (b GrayStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:gray_stained_glass_pane", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b GrayStainedGlassPane) New(props BlockProperties) Block {
	return GrayStainedGlassPane{
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
	}
}