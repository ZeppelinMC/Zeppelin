package block

import (
	"strconv"
)

type CopperTrapdoor struct {
	Open bool
	Powered bool
	Waterlogged bool
	Facing string
	Half string
}

func (b CopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:copper_trapdoor", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CopperTrapdoor) New(props BlockProperties) Block {
	return CopperTrapdoor{
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
	}
}