package block

import (
	"strconv"
)

type Fire struct {
	East bool
	North bool
	South bool
	Up bool
	West bool
	Age int
}

func (b Fire) Encode() (string, BlockProperties) {
	return "minecraft:fire", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"west": strconv.FormatBool(b.West),
		"age": strconv.Itoa(b.Age),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b Fire) New(props BlockProperties) Block {
	return Fire{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Up: props["up"] != "false",
		West: props["west"] != "false",
		Age: atoi(props["age"]),
		East: props["east"] != "false",
	}
}