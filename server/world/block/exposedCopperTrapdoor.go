package block

import (
	"strconv"
)

type ExposedCopperTrapdoor struct {
	Powered bool
	Waterlogged bool
	Facing string
	Half string
	Open bool
}

func (b ExposedCopperTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:exposed_copper_trapdoor", BlockProperties{
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b ExposedCopperTrapdoor) New(props BlockProperties) Block {
	return ExposedCopperTrapdoor{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}