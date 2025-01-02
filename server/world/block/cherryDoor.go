package block

import (
	"strconv"
)

type CherryDoor struct {
	Half string
	Hinge string
	Open bool
	Powered bool
	Facing string
}

func (b CherryDoor) Encode() (string, BlockProperties) {
	return "minecraft:cherry_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CherryDoor) New(props BlockProperties) Block {
	return CherryDoor{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
	}
}