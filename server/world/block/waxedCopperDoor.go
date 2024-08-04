package block

import (
	"strconv"
)

type WaxedCopperDoor struct {
	Open bool
	Powered bool
	Facing string
	Half string
	Hinge string
}

func (b WaxedCopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_copper_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WaxedCopperDoor) New(props BlockProperties) Block {
	return WaxedCopperDoor{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
	}
}