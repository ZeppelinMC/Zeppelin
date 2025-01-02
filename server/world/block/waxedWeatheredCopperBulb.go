package block

import (
	"strconv"
)

type WaxedWeatheredCopperBulb struct {
	Lit bool
	Powered bool
}

func (b WaxedWeatheredCopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_copper_bulb", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WaxedWeatheredCopperBulb) New(props BlockProperties) Block {
	return WaxedWeatheredCopperBulb{
		Lit: props["lit"] != "false",
		Powered: props["powered"] != "false",
	}
}