package block

import (
	"strconv"
)

type Rail struct {
	Shape string
	Waterlogged bool
}

func (b Rail) Encode() (string, BlockProperties) {
	return "minecraft:rail", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b Rail) New(props BlockProperties) Block {
	return Rail{
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}