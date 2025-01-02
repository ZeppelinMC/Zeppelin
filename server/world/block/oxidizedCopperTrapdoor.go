package block

import (
	"strconv"
)

type OxidizedCopperTrapdoor struct {
	Facing string
	Half string
	Open bool
	Powered bool
	Waterlogged bool
}

func (b OxidizedCopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_copper_trapdoor", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b OxidizedCopperTrapdoor) New(props BlockProperties) Block {
	return OxidizedCopperTrapdoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}