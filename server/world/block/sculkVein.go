package block

import (
	"strconv"
)

type SculkVein struct {
	North bool
	South bool
	Up bool
	Waterlogged bool
	West bool
	Down bool
	East bool
}

func (b SculkVein) Encode() (string, BlockProperties) {
	return "minecraft:sculk_vein", BlockProperties{
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"down": strconv.FormatBool(b.Down),
		"east": strconv.FormatBool(b.East),
	}
}

func (b SculkVein) New(props BlockProperties) Block {
	return SculkVein{
		Up: props["up"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		Down: props["down"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}