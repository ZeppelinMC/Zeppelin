package block

import (
	"strconv"
)

type BirchTrapdoor struct {
	Powered bool
	Waterlogged bool
	Facing string
	Half string
	Open bool
}

func (b BirchTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:birch_trapdoor", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b BirchTrapdoor) New(props BlockProperties) Block {
	return BirchTrapdoor{
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}