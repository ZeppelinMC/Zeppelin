package block

import (
	"strconv"
)

type MushroomStem struct {
	Up bool
	West bool
	Down bool
	East bool
	North bool
	South bool
}

func (b MushroomStem) Encode() (string, BlockProperties) {
	return "minecraft:mushroom_stem", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"up": strconv.FormatBool(b.Up),
		"west": strconv.FormatBool(b.West),
		"down": strconv.FormatBool(b.Down),
		"east": strconv.FormatBool(b.East),
	}
}

func (b MushroomStem) New(props BlockProperties) Block {
	return MushroomStem{
		Up: props["up"] != "false",
		West: props["west"] != "false",
		Down: props["down"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}