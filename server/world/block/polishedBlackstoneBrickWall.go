package block

import (
	"strconv"
)

type PolishedBlackstoneBrickWall struct {
	Waterlogged bool
	West string
	East string
	North string
	South string
	Up bool
}

func (b PolishedBlackstoneBrickWall) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_brick_wall", BlockProperties{
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
	}
}

func (b PolishedBlackstoneBrickWall) New(props BlockProperties) Block {
	return PolishedBlackstoneBrickWall{
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
		North: props["north"],
	}
}