package block

import (
	"strconv"
)

type SpruceButton struct {
	Face string
	Facing string
	Powered bool
}

func (b SpruceButton) Encode() (string, BlockProperties) {
	return "minecraft:spruce_button", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
		"face": b.Face,
	}
}

func (b SpruceButton) New(props BlockProperties) Block {
	return SpruceButton{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}