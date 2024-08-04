package block

import (
	"strconv"
)

type PolishedDeepslateWall struct {
	East string
	North string
	South string
	Up bool
	Waterlogged bool
	West string
}

func (b PolishedDeepslateWall) Encode() (string, BlockProperties) {
	return "minecraft:polished_deepslate_wall", BlockProperties{
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
	}
}

func (b PolishedDeepslateWall) New(props BlockProperties) Block {
	return PolishedDeepslateWall{
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
	}
}