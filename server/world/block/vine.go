package block

import (
	"strconv"
)

type Vine struct {
	North bool
	South bool
	Up bool
	West bool
	East bool
}

func (b Vine) Encode() (string, BlockProperties) {
	return "minecraft:vine", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"up": strconv.FormatBool(b.Up),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
	}
}

func (b Vine) New(props BlockProperties) Block {
	return Vine{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Up: props["up"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
	}
}