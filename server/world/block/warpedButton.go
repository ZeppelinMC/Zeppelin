package block

import (
	"strconv"
)

type WarpedButton struct {
	Facing string
	Powered bool
	Face string
}

func (b WarpedButton) Encode() (string, BlockProperties) {
	return "minecraft:warped_button", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
		"face": b.Face,
	}
}

func (b WarpedButton) New(props BlockProperties) Block {
	return WarpedButton{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
		Face: props["face"],
	}
}