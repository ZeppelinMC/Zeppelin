package block

import (
	"strconv"
)

type WeatheredCopperTrapdoor struct {
	Half string
	Open bool
	Powered bool
	Waterlogged bool
	Facing string
}

func (b WeatheredCopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:weathered_copper_trapdoor", BlockProperties{
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b WeatheredCopperTrapdoor) New(props BlockProperties) Block {
	return WeatheredCopperTrapdoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}