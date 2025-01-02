package block

import (
	"strconv"
)

type MossyStoneBrickWall struct {
	North string
	South string
	Up bool
	Waterlogged bool
	West string
	East string
}

func (b MossyStoneBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:mossy_stone_brick_wall", BlockProperties{
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
	}
}

func (b MossyStoneBrickWall) New(props BlockProperties) Block {
	return MossyStoneBrickWall{
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
		South: props["south"],
	}
}