package block

import (
	"strconv"
)

type PolishedBlackstoneButton struct {
	Face string
	Facing string
	Powered bool
}

func (b PolishedBlackstoneButton) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_button", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"face": b.Face,
		"facing": b.Facing,
	}
}

func (b PolishedBlackstoneButton) New(props BlockProperties) Block {
	return PolishedBlackstoneButton{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}