package block

import (
	"strconv"
)

type PolishedBlackstoneWall struct {
	Up bool
	Waterlogged bool
	West string
	East string
	North string
	South string
}

func (b PolishedBlackstoneWall) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_wall", BlockProperties{
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
	}
}

func (b PolishedBlackstoneWall) New(props BlockProperties) Block {
	return PolishedBlackstoneWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}