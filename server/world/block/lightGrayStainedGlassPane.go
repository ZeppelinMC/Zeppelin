package block

import (
	"strconv"
)

type LightGrayStainedGlassPane struct {
	South bool
	Waterlogged bool
	West bool
	East bool
	North bool
}

func (b LightGrayStainedGlassPane) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_stained_glass_pane", BlockProperties{
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
	}
}

func (b LightGrayStainedGlassPane) New(props BlockProperties) Block {
	return LightGrayStainedGlassPane{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
	}
}