package block

import (
	"strconv"
)

type SpruceLeaves struct {
	Persistent bool
	Waterlogged bool
	Distance int
}

func (b SpruceLeaves) Encode() (string, BlockProperties) {
	return "minecraft:spruce_leaves", BlockProperties{
		"persistent": strconv.FormatBool(b.Persistent),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"distance": strconv.Itoa(b.Distance),
	}
}

func (b SpruceLeaves) New(props BlockProperties) Block {
	return SpruceLeaves{
		Waterlogged: props["waterlogged"] != "false",
		Distance: atoi(props["distance"]),
		Persistent: props["persistent"] != "false",
	}
}