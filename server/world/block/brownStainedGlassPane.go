package block

import (
	"strconv"
)

type BrownStainedGlassPane struct {
	Waterlogged bool
	West bool
	East bool
	North bool
	South bool
}

func (b BrownStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:brown_stained_glass_pane", BlockProperties{
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BrownStainedGlassPane) New(props BlockProperties) Block {
	return BrownStainedGlassPane{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}