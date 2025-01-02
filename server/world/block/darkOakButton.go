package block

import (
	"strconv"
)

type DarkOakButton struct {
	Powered bool
	Face string
	Facing string
}

func (b DarkOakButton) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_button", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"face": b.Face,
		"facing": b.Facing,
	}
}

func (b DarkOakButton) New(props BlockProperties) Block {
	return DarkOakButton{
		Powered: props["powered"] != "false",
		Face: props["face"],
		Facing: props["facing"],
	}
}