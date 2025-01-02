package block

import (
	"strconv"
)

type BirchLeaves struct {
	Persistent bool
	Waterlogged bool
	Distance int
}

func (b BirchLeaves) Encode() (string, BlockProperties) {
	return "minecraft:birch_leaves", BlockProperties{
		"persistent": strconv.FormatBool(b.Persistent),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"distance": strconv.Itoa(b.Distance),
	}
}

func (b BirchLeaves) New(props BlockProperties) Block {
	return BirchLeaves{
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Distance: atoi(props["distance"]),
	}
}