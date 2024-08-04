package block

import (
	"strconv"
)

type WarpedTrapdoor struct {
	Half string
	Open bool
	Powered bool
	Waterlogged bool
	Facing string
}

func (b WarpedTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:warped_trapdoor", BlockProperties{
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b WarpedTrapdoor) New(props BlockProperties) Block {
	return WarpedTrapdoor{
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}