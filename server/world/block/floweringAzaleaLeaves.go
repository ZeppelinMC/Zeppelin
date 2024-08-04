package block

import (
	"strconv"
)

type FloweringAzaleaLeaves struct {
	Distance int
	Persistent bool
	Waterlogged bool
}

func (b FloweringAzaleaLeaves) Encode() (string, BlockProperties) {
	return "minecraft:flowering_azalea_leaves", BlockProperties{
		"distance": strconv.Itoa(b.Distance),
		"persistent": strconv.FormatBool(b.Persistent),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b FloweringAzaleaLeaves) New(props BlockProperties) Block {
	return FloweringAzaleaLeaves{
		Distance: atoi(props["distance"]),
		Persistent: props["persistent"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}