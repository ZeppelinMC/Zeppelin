package block

import (
	"strconv"
)

type CobbledDeepslateWall struct {
	South string
	Up bool
	Waterlogged bool
	West string
	East string
	North string
}

func (b CobbledDeepslateWall) Encode() (string, BlockProperties) {
	return "minecraft:cobbled_deepslate_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b CobbledDeepslateWall) New(props BlockProperties) Block {
	return CobbledDeepslateWall{
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
	}
}