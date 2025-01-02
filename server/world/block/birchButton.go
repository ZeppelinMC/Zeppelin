package block

import (
	"strconv"
)

type BirchButton struct {
	Face string
	Facing string
	Powered bool
}

func (b BirchButton) Encode() (string, BlockProperties) {
	return "minecraft:birch_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b BirchButton) New(props BlockProperties) Block {
	return BirchButton{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}