package block

import (
	"strconv"
)

type AcaciaButton struct {
	Facing string
	Powered bool
	Face string
}

func (b AcaciaButton) Encode() (string, BlockProperties) {
	return "minecraft:acacia_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b AcaciaButton) New(props BlockProperties) Block {
	return AcaciaButton{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}