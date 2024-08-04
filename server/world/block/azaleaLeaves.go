package block

import (
	"strconv"
)

type AzaleaLeaves struct {
	Persistent bool
	Waterlogged bool
	Distance int
}

func (b AzaleaLeaves) Encode() (string, BlockProperties) {
	return "minecraft:azalea_leaves", BlockProperties{
		"persistent": strconv.FormatBool(b.Persistent),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"distance": strconv.Itoa(b.Distance),
	}
}

func (b AzaleaLeaves) New(props BlockProperties) Block {
	return AzaleaLeaves{
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Distance: atoi(props["distance"]),
	}
}