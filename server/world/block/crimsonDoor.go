package block

import (
	"strconv"
)

type CrimsonDoor struct {
	Half string
	Hinge string
	Open bool
	Powered bool
	Facing string
}

func (b CrimsonDoor) Encode() (string, BlockProperties) {
	return "minecraft:crimson_door", BlockProperties{
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b CrimsonDoor) New(props BlockProperties) Block {
	return CrimsonDoor{
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}