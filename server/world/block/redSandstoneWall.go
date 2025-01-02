package block

import (
	"strconv"
)

type RedSandstoneWall struct {
	Up bool
	Waterlogged bool
	West string
	East string
	North string
	South string
}

func (b RedSandstoneWall) Encode() (string, BlockProperties) {
	return "minecraft:red_sandstone_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b RedSandstoneWall) New(props BlockProperties) Block {
	return RedSandstoneWall{
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
	}
}