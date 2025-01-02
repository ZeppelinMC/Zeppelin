package block

import (
	"strconv"
)

type ExposedCopperDoor struct {
	Facing string
	Half string
	Hinge string
	Open bool
	Powered bool
}

func (b ExposedCopperDoor) Encode() (string, BlockProperties) {
	return "minecraft:exposed_copper_door", BlockProperties{
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
	}
}

func (b ExposedCopperDoor) New(props BlockProperties) Block {
	return ExposedCopperDoor{
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}