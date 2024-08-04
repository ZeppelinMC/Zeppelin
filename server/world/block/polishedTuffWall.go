package block

import (
	"strconv"
)

type PolishedTuffWall struct {
	East string
	North string
	South string
	Up bool
	Waterlogged bool
	West string
}

func (b PolishedTuffWall) Encode() (string, BlockProperties) {
	return "minecraft:polished_tuff_wall", BlockProperties{
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
	}
}

func (b PolishedTuffWall) New(props BlockProperties) Block {
	return PolishedTuffWall{
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
	}
}