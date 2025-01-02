package block

import (
	"strconv"
)

type MangroveTrapdoor struct {
	Powered bool
	Waterlogged bool
	Facing string
	Half string
	Open bool
}

func (b MangroveTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_trapdoor", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MangroveTrapdoor) New(props BlockProperties) Block {
	return MangroveTrapdoor{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}