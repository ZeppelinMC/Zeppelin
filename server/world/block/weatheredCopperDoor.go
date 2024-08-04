package block

import (
	"strconv"
)

type WeatheredCopperDoor struct {
	Powered bool
	Facing string
	Half string
	Hinge string
	Open bool
}

func (b WeatheredCopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:weathered_copper_door", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b WeatheredCopperDoor) New(props BlockProperties) Block {
	return WeatheredCopperDoor{
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}