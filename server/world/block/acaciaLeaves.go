package block

import (
	"strconv"
)

type AcaciaLeaves struct {
	Waterlogged bool
	Distance int
	Persistent bool
}

func (b AcaciaLeaves) Encode() (string, BlockProperties) {
	return "minecraft:acacia_leaves", BlockProperties{
		"distance": strconv.Itoa(b.Distance),
		"persistent": strconv.FormatBool(b.Persistent),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AcaciaLeaves) New(props BlockProperties) Block {
	return AcaciaLeaves{
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Distance: atoi(props["distance"]),
	}
}