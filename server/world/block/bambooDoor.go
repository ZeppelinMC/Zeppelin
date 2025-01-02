package block

import (
	"strconv"
)

type BambooDoor struct {
	Powered bool
	Facing string
	Half string
	Hinge string
	Open bool
}

func (b BambooDoor) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_door", BlockProperties{
		"half": b.Half,
		"hinge": b.Hinge,
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
	}
}

func (b BambooDoor) New(props BlockProperties) Block {
	return BambooDoor{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Half: props["half"],
		Hinge: props["hinge"],
	}
}