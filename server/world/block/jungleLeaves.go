package block

import (
	"strconv"
)

type JungleLeaves struct {
	Distance int
	Persistent bool
	Waterlogged bool
}

func (b JungleLeaves) Encode() (string, BlockProperties) {
	return "minecraft:jungle_leaves", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"distance": strconv.Itoa(b.Distance),
		"persistent": strconv.FormatBool(b.Persistent),
	}
}

func (b JungleLeaves) New(props BlockProperties) Block {
	return JungleLeaves{
		Distance: atoi(props["distance"]),
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}