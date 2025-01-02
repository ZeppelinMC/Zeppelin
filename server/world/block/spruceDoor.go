package block

import (
	"strconv"
)

type SpruceDoor struct {
	Hinge string
	Open bool
	Powered bool
	Facing string
	Half string
}

func (b SpruceDoor) Encode() (string, BlockProperties) {
	return "minecraft:spruce_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b SpruceDoor) New(props BlockProperties) Block {
	return SpruceDoor{
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}