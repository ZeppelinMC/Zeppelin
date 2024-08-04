package block

import (
	"strconv"
)

type StoneBrickWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b StoneBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:stone_brick_wall", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
	}
}

func (b StoneBrickWall) New(props BlockProperties) Block {
	return StoneBrickWall{
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
	}
}