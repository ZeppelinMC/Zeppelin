package block

import (
	"strconv"
)

type JungleDoor struct {
	Half string
	Hinge string
	Open bool
	Powered bool
	Facing string
}

func (b JungleDoor) Encode() (string, BlockProperties) {
	return "minecraft:jungle_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b JungleDoor) New(props BlockProperties) Block {
	return JungleDoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
	}
}