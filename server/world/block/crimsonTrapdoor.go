package block

import (
	"strconv"
)

type CrimsonTrapdoor struct {
	Half string
	Open bool
	Powered bool
	Waterlogged bool
	Facing string
}

func (b CrimsonTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:crimson_trapdoor", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CrimsonTrapdoor) New(props BlockProperties) Block {
	return CrimsonTrapdoor{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}