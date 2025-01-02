package block

import (
	"strconv"
)

type OrangeStainedGlassPane struct {
	South bool
	Waterlogged bool
	West bool
	East bool
	North bool
}

func (b OrangeStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:orange_stained_glass_pane", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b OrangeStainedGlassPane) New(props BlockProperties) Block {
	return OrangeStainedGlassPane{
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
	}
}