package block

import (
	"strconv"
)

type WhiteStainedGlassPane struct {
	North bool
	South bool
	Waterlogged bool
	West bool
	East bool
}

func (b WhiteStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:white_stained_glass_pane", BlockProperties{
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
	}
}

func (b WhiteStainedGlassPane) New(props BlockProperties) Block {
	return WhiteStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}