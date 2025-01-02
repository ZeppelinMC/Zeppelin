package block

import (
	"strconv"
)

type BlackstoneWall struct {
	Waterlogged bool
	West string
	East string
	North string
	South string
	Up bool
}

func (b BlackstoneWall) Encode() (string, BlockProperties) {
	return "minecraft:blackstone_wall", BlockProperties{
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
	}
}

func (b BlackstoneWall) New(props BlockProperties) Block {
	return BlackstoneWall{
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
	}
}