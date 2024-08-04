package block

import (
	"strconv"
)

type Lever struct {
	Face string
	Facing string
	Powered bool
}

func (b Lever) Encode() (string, BlockProperties) {
	return "minecraft:lever", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b Lever) New(props BlockProperties) Block {
	return Lever{
		Face: props["face"],
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}