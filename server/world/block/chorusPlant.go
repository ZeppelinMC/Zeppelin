package block

import (
	"strconv"
)

type ChorusPlant struct {
	Up bool
	West bool
	Down bool
	East bool
	North bool
	South bool
}

func (b ChorusPlant) Encode() (string, BlockProperties) {
	return "minecraft:chorus_plant", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"up": strconv.FormatBool(b.Up),
		"west": strconv.FormatBool(b.West),
		"down": strconv.FormatBool(b.Down),
		"east": strconv.FormatBool(b.East),
	}
}

func (b ChorusPlant) New(props BlockProperties) Block {
	return ChorusPlant{
		Down: props["down"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Up: props["up"] != "false",
		West: props["west"] != "false",
	}
}