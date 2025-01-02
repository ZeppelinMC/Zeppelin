package block

import (
	"strconv"
)

type Observer struct {
	Facing string
	Powered bool
}

func (b Observer) Encode() (string, BlockProperties) {
	return "minecraft:observer", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
	}
}

func (b Observer) New(props BlockProperties) Block {
	return Observer{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}