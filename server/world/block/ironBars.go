package block

import (
	"strconv"
)

type IronBars struct {
	East bool
	North bool
	South bool
	Waterlogged bool
	West bool
}

func (b IronBars) Encode() (string, BlockProperties) {
	return "minecraft:iron_bars", BlockProperties{
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
	}
}

func (b IronBars) New(props BlockProperties) Block {
	return IronBars{
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}