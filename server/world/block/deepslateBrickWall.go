package block

import (
	"strconv"
)

type DeepslateBrickWall struct {
	Up bool
	Waterlogged bool
	West string
	East string
	North string
	South string
}

func (b DeepslateBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_brick_wall", BlockProperties{
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DeepslateBrickWall) New(props BlockProperties) Block {
	return DeepslateBrickWall{
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
	}
}