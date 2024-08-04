package block

import (
	"strconv"
)

type GlowLichen struct {
	East bool
	North bool
	South bool
	Up bool
	Waterlogged bool
	West bool
	Down bool
}

func (b GlowLichen) Encode() (string, BlockProperties) {
	return "minecraft:glow_lichen", BlockProperties{
		"up": strconv.FormatBool(b.Up),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"down": strconv.FormatBool(b.Down),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"south": strconv.FormatBool(b.South),
	}
}

func (b GlowLichen) New(props BlockProperties) Block {
	return GlowLichen{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		Down: props["down"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
		Up: props["up"] != "false",
	}
}