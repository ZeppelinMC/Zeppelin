package block

import (
	"strconv"
)

type WaxedExposedCopperDoor struct {
	Facing string
	Half string
	Hinge string
	Open bool
	Powered bool
}

func (b WaxedExposedCopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_copper_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WaxedExposedCopperDoor) New(props BlockProperties) Block {
	return WaxedExposedCopperDoor{
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
	}
}