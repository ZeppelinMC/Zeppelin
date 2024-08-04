package block

import (
	"strconv"
)

type TuffBrickWall struct {
	North string
	South string
	Up bool
	Waterlogged bool
	West string
	East string
}

func (b TuffBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:tuff_brick_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b TuffBrickWall) New(props BlockProperties) Block {
	return TuffBrickWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}