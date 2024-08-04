package block

import (
	"strconv"
)

type DeepslateTileWall struct {
	West string
	East string
	North string
	South string
	Up bool
	Waterlogged bool
}

func (b DeepslateTileWall) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_tile_wall", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": b.West,
		"east": b.East,
		"north": b.North,
		"south": b.South,
		"up": strconv.FormatBool(b.Up),
	}
}

func (b DeepslateTileWall) New(props BlockProperties) Block {
	return DeepslateTileWall{
		North: props["north"],
		South: props["south"],
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"],
		East: props["east"],
	}
}