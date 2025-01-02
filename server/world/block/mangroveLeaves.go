package block

import (
	"strconv"
)

type MangroveLeaves struct {
	Distance int
	Persistent bool
	Waterlogged bool
}

func (b MangroveLeaves) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_leaves", BlockProperties{
		"distance": strconv.Itoa(b.Distance),
		"persistent": strconv.FormatBool(b.Persistent),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MangroveLeaves) New(props BlockProperties) Block {
	return MangroveLeaves{
		Distance: atoi(props["distance"]),
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}