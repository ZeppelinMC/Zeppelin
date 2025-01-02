package block

import (
	"strconv"
)

type DarkOakLeaves struct {
	Distance int
	Persistent bool
	Waterlogged bool
}

func (b DarkOakLeaves) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_leaves", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"distance": strconv.Itoa(b.Distance),
		"persistent": strconv.FormatBool(b.Persistent),
	}
}

func (b DarkOakLeaves) New(props BlockProperties) Block {
	return DarkOakLeaves{
		Distance: atoi(props["distance"]),
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}