package block

import (
	"strconv"
)

type CopperDoor struct {
	Open bool
	Powered bool
	Facing string
	Half string
	Hinge string
}

func (b CopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:copper_door", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b CopperDoor) New(props BlockProperties) Block {
	return CopperDoor{
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
	}
}