package block

import (
	"strconv"
)

type PoweredRail struct {
	Powered bool
	Shape string
	Waterlogged bool
}

func (b PoweredRail) Encode() (string, BlockProperties) {
	return "minecraft:powered_rail", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PoweredRail) New(props BlockProperties) Block {
	return PoweredRail{
		Powered: props["powered"] != "false",
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}