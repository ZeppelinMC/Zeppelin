package block

import (
	"strconv"
)

type WaxedWeatheredCopperDoor struct {
	Hinge string
	Open bool
	Powered bool
	Facing string
	Half string
}

func (b WaxedWeatheredCopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_copper_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WaxedWeatheredCopperDoor) New(props BlockProperties) Block {
	return WaxedWeatheredCopperDoor{
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}