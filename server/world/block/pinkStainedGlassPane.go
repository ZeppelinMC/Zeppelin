package block

import (
	"strconv"
)

type PinkStainedGlassPane struct {
	West bool
	East bool
	North bool
	South bool
	Waterlogged bool
}

func (b PinkStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:pink_stained_glass_pane", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b PinkStainedGlassPane) New(props BlockProperties) Block {
	return PinkStainedGlassPane{
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
	}
}