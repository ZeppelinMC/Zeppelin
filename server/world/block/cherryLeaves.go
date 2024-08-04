package block

import (
	"strconv"
)

type CherryLeaves struct {
	Distance int
	Persistent bool
	Waterlogged bool
}

func (b CherryLeaves) Encode() (string, BlockProperties) {
	return "minecraft:cherry_leaves", BlockProperties{
		"persistent": strconv.FormatBool(b.Persistent),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"distance": strconv.Itoa(b.Distance),
	}
}

func (b CherryLeaves) New(props BlockProperties) Block {
	return CherryLeaves{
		Distance: atoi(props["distance"]),
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}