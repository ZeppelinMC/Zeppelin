package block

import (
	"strconv"
)

type MudBrickWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b MudBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:mud_brick_wall", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
	}
}

func (b MudBrickWall) New(props BlockProperties) Block {
	return MudBrickWall{
		East: props["east"],
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
	}
}