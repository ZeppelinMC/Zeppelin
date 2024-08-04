package block

import (
	"strconv"
)

type OakTrapdoor struct {
	Waterlogged bool
	Facing string
	Half string
	Open bool
	Powered bool
}

func (b OakTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:oak_trapdoor", BlockProperties{
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b OakTrapdoor) New(props BlockProperties) Block {
	return OakTrapdoor{
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}