package block

import (
	"strconv"
)

type DarkOakTrapdoor struct {
	Half string
	Open bool
	Powered bool
	Waterlogged bool
	Facing string
}

func (b DarkOakTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_trapdoor", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b DarkOakTrapdoor) New(props BlockProperties) Block {
	return DarkOakTrapdoor{
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}