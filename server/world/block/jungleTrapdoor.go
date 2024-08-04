package block

import (
	"strconv"
)

type JungleTrapdoor struct {
	Facing string
	Half string
	Open bool
	Powered bool
	Waterlogged bool
}

func (b JungleTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:jungle_trapdoor", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b JungleTrapdoor) New(props BlockProperties) Block {
	return JungleTrapdoor{
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
	}
}