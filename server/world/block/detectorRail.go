package block

import (
	"strconv"
)

type DetectorRail struct {
	Powered bool
	Shape string
	Waterlogged bool
}

func (b DetectorRail) Encode() (string, BlockProperties) {
	return "minecraft:detector_rail", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"shape": b.Shape,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b DetectorRail) New(props BlockProperties) Block {
	return DetectorRail{
		Powered: props["powered"] != "false",
		Shape: props["shape"],
		Waterlogged: props["waterlogged"] != "false",
	}
}