package block

import (
	"strconv"
)

type IronTrapdoor struct {
	Half string
	Open bool
	Powered bool
	Waterlogged bool
	Facing string
}

func (b IronTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:iron_trapdoor", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b IronTrapdoor) New(props BlockProperties) Block {
	return IronTrapdoor{
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}