package block

import (
	"strconv"
)

type WaxedWeatheredCopperTrapdoor struct {
	Waterlogged bool
	Facing string
	Half string
	Open bool
	Powered bool
}

func (b WaxedWeatheredCopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_copper_trapdoor", BlockProperties{
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b WaxedWeatheredCopperTrapdoor) New(props BlockProperties) Block {
	return WaxedWeatheredCopperTrapdoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}