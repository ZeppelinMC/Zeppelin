package block

import (
	"strconv"
)

type CrimsonButton struct {
	Face string
	Facing string
	Powered bool
}

func (b CrimsonButton) Encode() (string, BlockProperties) {
	return "minecraft:crimson_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CrimsonButton) New(props BlockProperties) Block {
	return CrimsonButton{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}