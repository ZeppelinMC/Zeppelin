package block

import (
	"strconv"
)

type RedNetherBrickWall struct {
	Waterlogged bool
	West string
	East string
	North string
	South string
	Up bool
}

func (b RedNetherBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:red_nether_brick_wall", BlockProperties{
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
	}
}

func (b RedNetherBrickWall) New(props BlockProperties) Block {
	return RedNetherBrickWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}