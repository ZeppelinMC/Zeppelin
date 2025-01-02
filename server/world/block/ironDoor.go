package block

import (
	"strconv"
)

type IronDoor struct {
	Half string
	Hinge string
	Open bool
	Powered bool
	Facing string
}

func (b IronDoor) Encode() (string, BlockProperties) {
	return "minecraft:iron_door", BlockProperties{
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
	}
}

func (b IronDoor) New(props BlockProperties) Block {
	return IronDoor{
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}