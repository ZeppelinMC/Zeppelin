package block

import (
	"strconv"
)

type WarpedDoor struct {
	Powered bool
	Facing string
	Half string
	Hinge string
	Open bool
}

func (b WarpedDoor) Encode() (string, BlockProperties) {
	return "minecraft:warped_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WarpedDoor) New(props BlockProperties) Block {
	return WarpedDoor{
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}