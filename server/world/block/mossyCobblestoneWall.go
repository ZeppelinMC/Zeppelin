package block

import (
	"strconv"
)

type MossyCobblestoneWall struct {
	South string
	Up bool
	Waterlogged bool
	West string
	East string
	North string
}

func (b MossyCobblestoneWall) Encode() (string, BlockProperties) {
	return "minecraft:mossy_cobblestone_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b MossyCobblestoneWall) New(props BlockProperties) Block {
	return MossyCobblestoneWall{
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
	}
}