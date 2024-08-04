package block

import (
	"strconv"
)

type BirchDoor struct {
	Facing string
	Half string
	Hinge string
	Open bool
	Powered bool
}

func (b BirchDoor) Encode() (string, BlockProperties) {
	return "minecraft:birch_door", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b BirchDoor) New(props BlockProperties) Block {
	return BirchDoor{
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}