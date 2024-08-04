package block

import (
	"strconv"
)

type WaxedOxidizedCopperDoor struct {
	Hinge string
	Open bool
	Powered bool
	Facing string
	Half string
}

func (b WaxedOxidizedCopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_copper_door", BlockProperties{
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b WaxedOxidizedCopperDoor) New(props BlockProperties) Block {
	return WaxedOxidizedCopperDoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
	}
}