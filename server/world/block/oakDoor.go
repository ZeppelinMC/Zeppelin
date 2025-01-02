package block

import (
	"strconv"
)

type OakDoor struct {
	Powered bool
	Facing string
	Half string
	Hinge string
	Open bool
}

func (b OakDoor) Encode() (string, BlockProperties) {
	return "minecraft:oak_door", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b OakDoor) New(props BlockProperties) Block {
	return OakDoor{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
	}
}