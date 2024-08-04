package block

import (
	"strconv"
)

type Scaffolding struct {
	Waterlogged bool
	Bottom bool
	Distance int
}

func (b Scaffolding) Encode() (string, BlockProperties) {
	return "minecraft:scaffolding", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"bottom": strconv.FormatBool(b.Bottom),
		"distance": strconv.Itoa(b.Distance),
	}
}

func (b Scaffolding) New(props BlockProperties) Block {
	return Scaffolding{
		Waterlogged: props["waterlogged"] != "false",
		Bottom: props["bottom"] != "false",
		Distance: atoi(props["distance"]),
	}
}