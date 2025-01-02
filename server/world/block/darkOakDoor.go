package block

import (
	"strconv"
)

type DarkOakDoor struct {
	Facing string
	Half string
	Hinge string
	Open bool
	Powered bool
}

func (b DarkOakDoor) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_door", BlockProperties{
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b DarkOakDoor) New(props BlockProperties) Block {
	return DarkOakDoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
	}
}