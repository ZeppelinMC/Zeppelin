package block

import (
	"strconv"
)

type CherryTrapdoor struct {
	Facing string
	Half string
	Open bool
	Powered bool
	Waterlogged bool
}

func (b CherryTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:cherry_trapdoor", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CherryTrapdoor) New(props BlockProperties) Block {
	return CherryTrapdoor{
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
	}
}