package block

import (
	"strconv"
)

type OakLeaves struct {
	Persistent bool
	Waterlogged bool
	Distance int
}

func (b OakLeaves) Encode() (string, BlockProperties) {
	return "minecraft:oak_leaves", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"distance": strconv.Itoa(b.Distance),
		"persistent": strconv.FormatBool(b.Persistent),
	}
}

func (b OakLeaves) New(props BlockProperties) Block {
	return OakLeaves{
		Distance: atoi(props["distance"]),
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}