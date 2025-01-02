package block

import (
	"strconv"
)

type JungleButton struct {
	Face string
	Facing string
	Powered bool
}

func (b JungleButton) Encode() (string, BlockProperties) {
	return "minecraft:jungle_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b JungleButton) New(props BlockProperties) Block {
	return JungleButton{
		Powered: props["powered"] != "false",
		Face: props["face"],
		Facing: props["facing"],
	}
}