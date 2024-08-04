package block

import (
	"strconv"
)

type WaxedExposedCopperBulb struct {
	Lit bool
	Powered bool
}

func (b WaxedExposedCopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:waxed_exposed_copper_bulb", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WaxedExposedCopperBulb) New(props BlockProperties) Block {
	return WaxedExposedCopperBulb{
		Lit: props["lit"] != "false",
		Powered: props["powered"] != "false",
	}
}