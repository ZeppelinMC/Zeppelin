package block

import (
	"strconv"
)

type AcaciaTrapdoor struct {
	Waterlogged bool
	Facing string
	Half string
	Open bool
	Powered bool
}

func (b AcaciaTrapdoor) Encode() (string, BlockProperties) {
	return "minecraft:acacia_trapdoor", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"half": b.Half,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b AcaciaTrapdoor) New(props BlockProperties) Block {
	return AcaciaTrapdoor{
		Half: props["half"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}