package block

import (
	"strconv"
)

type CobblestoneWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b CobblestoneWall) Encode() (string, BlockProperties) {
	return "minecraft:cobblestone_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b CobblestoneWall) New(props BlockProperties) Block {
	return CobblestoneWall{
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
	}
}