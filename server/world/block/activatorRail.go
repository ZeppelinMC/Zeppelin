package block

import (
	"strconv"
)

type ActivatorRail struct {
	Powered bool
	Shape string
	Waterlogged bool
}

func (b ActivatorRail) Encode() (string, BlockProperties) {
	return "minecraft:activator_rail", BlockProperties{
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b ActivatorRail) New(props BlockProperties) Block {
	return ActivatorRail{
		Powered: props["powered"] != "false",
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}