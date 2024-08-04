package block

import (
	"strconv"
)

type RedstoneWire struct {
	East string
	North string
	Power int
	South string
	West string
}

func (b RedstoneWire) Encode() (string, BlockProperties) {
	return "minecraft:redstone_wire", BlockProperties{
		"north": b.North,
		"power": strconv.Itoa(b.Power),
		"south": b.South,
		"west": b.West,
		"east": b.East,
	}
}

func (b RedstoneWire) New(props BlockProperties) Block {
	return RedstoneWire{
		Power: atoi(props["power"]),
		South: props["south"],
		West: props["west"],
		East: props["east"],
		North: props["north"],
	}
}