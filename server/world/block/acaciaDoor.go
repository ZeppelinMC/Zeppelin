package block

import (
	"strconv"
)

type AcaciaDoor struct {
	Open bool
	Powered bool
	Facing string
	Half string
	Hinge string
}

func (b AcaciaDoor) Encode() (string, BlockProperties) {
	return "minecraft:acacia_door", BlockProperties{
		"facing": b.Facing,
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b AcaciaDoor) New(props BlockProperties) Block {
	return AcaciaDoor{
		Hinge: props["hinge"],
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
	}
}