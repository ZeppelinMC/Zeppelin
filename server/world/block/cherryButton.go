package block

import (
	"strconv"
)

type CherryButton struct {
	Facing string
	Powered bool
	Face string
}

func (b CherryButton) Encode() (string, BlockProperties) {
	return "minecraft:cherry_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CherryButton) New(props BlockProperties) Block {
	return CherryButton{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
		Face: props["face"],
	}
}