package block

import (
	"strconv"
)

type TuffWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b TuffWall) Encode() (string, BlockProperties) {
	return "minecraft:tuff_wall", BlockProperties{
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
	}
}

func (b TuffWall) New(props BlockProperties) Block {
	return TuffWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}