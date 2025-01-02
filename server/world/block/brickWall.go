package block

import (
	"strconv"
)

type BrickWall struct {
	Waterlogged bool
	West string
	East string
	North string
	South string
	Up bool
}

func (b BrickWall) Encode() (string, BlockProperties) {
	return "minecraft:brick_wall", BlockProperties{
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
	}
}

func (b BrickWall) New(props BlockProperties) Block {
	return BrickWall{
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
		South: props["south"],
	}
}