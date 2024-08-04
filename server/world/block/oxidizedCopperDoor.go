package block

import (
	"strconv"
)

type OxidizedCopperDoor struct {
	Powered bool
	Facing string
	Half string
	Hinge string
	Open bool
}

func (b OxidizedCopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_copper_door", BlockProperties{
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
	}
}

func (b OxidizedCopperDoor) New(props BlockProperties) Block {
	return OxidizedCopperDoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
	}
}