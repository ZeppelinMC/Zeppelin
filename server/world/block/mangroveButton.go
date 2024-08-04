package block

import (
	"strconv"
)

type MangroveButton struct {
	Facing string
	Powered bool
	Face string
}

func (b MangroveButton) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b MangroveButton) New(props BlockProperties) Block {
	return MangroveButton{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}