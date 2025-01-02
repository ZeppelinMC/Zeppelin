package block

import (
	"strconv"
)

type BambooTrapdoor struct {
	Facing string
	Half string
	Open bool
	Powered bool
	Waterlogged bool
}

func (b BambooTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_trapdoor", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BambooTrapdoor) New(props BlockProperties) Block {
	return BambooTrapdoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}