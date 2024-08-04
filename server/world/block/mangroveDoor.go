package block

import (
	"strconv"
)

type MangroveDoor struct {
	Powered bool
	Facing string
	Half string
	Hinge string
	Open bool
}

func (b MangroveDoor) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_door", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b MangroveDoor) New(props BlockProperties) Block {
	return MangroveDoor{
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
	}
}