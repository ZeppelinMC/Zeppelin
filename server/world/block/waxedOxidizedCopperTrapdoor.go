package block

import (
	"strconv"
)

type WaxedOxidizedCopperTrapdoor struct {
	Facing string
	Half string
	Open bool
	Powered bool
	Waterlogged bool
}

func (b WaxedOxidizedCopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_copper_trapdoor", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b WaxedOxidizedCopperTrapdoor) New(props BlockProperties) Block {
	return WaxedOxidizedCopperTrapdoor{
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}