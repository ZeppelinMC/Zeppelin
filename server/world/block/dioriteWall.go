package block

import (
	"strconv"
)

type DioriteWall struct {
	East string
	North string
	South string
	Up bool
	Waterlogged bool
	West string
}

func (b DioriteWall) Encode() (string, BlockProperties) {
	return "minecraft:diorite_wall", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
	}
}

func (b DioriteWall) New(props BlockProperties) Block {
	return DioriteWall{
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
		South: props["south"],
	}
}