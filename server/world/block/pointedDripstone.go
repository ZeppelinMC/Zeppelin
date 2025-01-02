package block

import (
	"strconv"
)

type PointedDripstone struct {
	VerticalDirection string
	Waterlogged bool
	Thickness string
}

func (b PointedDripstone) Encode() (string, BlockProperties) {
	return "minecraft:pointed_dripstone", BlockProperties{
		"thickness": b.Thickness,
		"vertical_direction": b.VerticalDirection,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PointedDripstone) New(props BlockProperties) Block {
	return PointedDripstone{
		VerticalDirection: props["vertical_direction"],
		Waterlogged: props["waterlogged"] != "false",
		Thickness: props["thickness"],
	}
}