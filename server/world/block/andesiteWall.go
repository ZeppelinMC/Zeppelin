package block

import (
	"strconv"
)

type AndesiteWall struct {
	North string
	South string
	Up bool
	Waterlogged bool
	West string
	East string
}

func (b AndesiteWall) Encode() (string, BlockProperties) {
	return "minecraft:andesite_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b AndesiteWall) New(props BlockProperties) Block {
	return AndesiteWall{
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
	}
}