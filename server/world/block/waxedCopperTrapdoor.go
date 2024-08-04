package block

import (
	"strconv"
)

type WaxedCopperTrapdoor struct {
	Waterlogged bool
	Facing string
	Half string
	Open bool
	Powered bool
}

func (b WaxedCopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_copper_trapdoor", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b WaxedCopperTrapdoor) New(props BlockProperties) Block {
	return WaxedCopperTrapdoor{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}