package block

import (
	"strconv"
)

type GraniteWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b GraniteWall) Encode() (string, BlockProperties) {
	return "minecraft:granite_wall", BlockProperties{
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b GraniteWall) New(props BlockProperties) Block {
	return GraniteWall{
		West: props["west"],
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}