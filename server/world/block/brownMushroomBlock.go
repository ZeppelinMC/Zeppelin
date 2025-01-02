package block

import (
	"strconv"
)

type BrownMushroomBlock struct {
	West bool
	Down bool
	East bool
	North bool
	South bool
	Up bool
}

func (b BrownMushroomBlock) Encode() (string, BlockProperties) {
	return "minecraft:brown_mushroom_block", BlockProperties{
		"down": strconv.FormatBool(b.Down),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"up": strconv.FormatBool(b.Up),
		"west": strconv.FormatBool(b.West),
	}
}

func (b BrownMushroomBlock) New(props BlockProperties) Block {
	return BrownMushroomBlock{
		Up: props["up"] != "false",
		West: props["west"] != "false",
		Down: props["down"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}