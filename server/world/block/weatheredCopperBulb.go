package block

import (
	"strconv"
)

type WeatheredCopperBulb struct {
	Lit bool
	Powered bool
}

func (b WeatheredCopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:weathered_copper_bulb", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WeatheredCopperBulb) New(props BlockProperties) Block {
	return WeatheredCopperBulb{
		Lit: props["lit"] != "false",
		Powered: props["powered"] != "false",
	}
}