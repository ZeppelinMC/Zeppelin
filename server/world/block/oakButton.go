package block

import (
	"strconv"
)

type OakButton struct {
	Facing string
	Powered bool
	Face string
}

func (b OakButton) Encode() (string, BlockProperties) {
	return "minecraft:oak_button", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
		"face": b.Face,
	}
}

func (b OakButton) New(props BlockProperties) Block {
	return OakButton{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}