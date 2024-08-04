package block

import (
	"strconv"
)

type WaxedExposedCopperTrapdoor struct {
	Facing string
	Half string
	Open bool
	Powered bool
	Waterlogged bool
}

func (b WaxedExposedCopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_copper_trapdoor", BlockProperties{
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b WaxedExposedCopperTrapdoor) New(props BlockProperties) Block {
	return WaxedExposedCopperTrapdoor{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}