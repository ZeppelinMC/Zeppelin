package block

import (
	"strconv"
)

type BambooButton struct {
	Face string
	Facing string
	Powered bool
}

func (b BambooButton) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b BambooButton) New(props BlockProperties) Block {
	return BambooButton{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
		Face: props["face"],
	}
}