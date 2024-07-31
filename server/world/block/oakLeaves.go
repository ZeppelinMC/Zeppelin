package block

import (
	"strconv"
)

type OakLeaves struct {
	Distance    int
	Persistent  bool
	Waterlogged bool
}

func (g OakLeaves) Encode() (string, BlockProperties) {
	return "minecraft:oak_leaves", BlockProperties{
		"distance":    strconv.Itoa(g.Distance),
		"persistent":  strconv.FormatBool(g.Persistent),
		"waterlogged": strconv.FormatBool(g.Waterlogged),
	}
}

func (g OakLeaves) New(props BlockProperties) Block {
	return OakLeaves{
		Distance:    atoi(props["distance"]),
		Persistent:  props["persistent"] == "true",
		Waterlogged: props["waterlogged"] == "true",
	}
}

var _ Block = (*OakLeaves)(nil)
