package block

import (
	"strconv"
)

type EndStoneBrickWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b EndStoneBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:end_stone_brick_wall", BlockProperties{
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
	}
}

func (b EndStoneBrickWall) New(props BlockProperties) Block {
	return EndStoneBrickWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}