package block

import (
	"strconv"
)

type NetherBrickWall struct {
	East string
	North string
	South string
	Up bool
	Waterlogged bool
	West string
}

func (b NetherBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:nether_brick_wall", BlockProperties{
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
	}
}

func (b NetherBrickWall) New(props BlockProperties) Block {
	return NetherBrickWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}