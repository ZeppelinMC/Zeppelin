package block

import (
	"strconv"
)

type PrismarineWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b PrismarineWall) Encode() (string, BlockProperties) {
	return "minecraft:prismarine_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b PrismarineWall) New(props BlockProperties) Block {
	return PrismarineWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}