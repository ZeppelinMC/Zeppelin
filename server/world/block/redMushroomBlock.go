package block

import (
	"strconv"
)

type RedMushroomBlock struct {
	Down bool
	East bool
	North bool
	South bool
	Up bool
	West bool
}

func (b RedMushroomBlock) Encode() (string, BlockProperties) {
	return "minecraft:red_mushroom_block", BlockProperties{
		"down": strconv.FormatBool(b.Down),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"up": strconv.FormatBool(b.Up),
		"west": strconv.FormatBool(b.West),
	}
}

func (b RedMushroomBlock) New(props BlockProperties) Block {
	return RedMushroomBlock{
		North: props["north"] != "false",
		South: props["south"] != "false",
		Up: props["up"] != "false",
		West: props["west"] != "false",
		Down: props["down"] != "false",
		East: props["east"] != "false",
	}
}