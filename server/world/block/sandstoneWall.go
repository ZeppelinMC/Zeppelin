package block

import (
	"strconv"
)

type SandstoneWall struct {
	South string
	Up bool
	Waterlogged bool
	West string
	East string
	North string
}

func (b SandstoneWall) Encode() (string, BlockProperties) {
	return "minecraft:sandstone_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b SandstoneWall) New(props BlockProperties) Block {
	return SandstoneWall{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
	}
}