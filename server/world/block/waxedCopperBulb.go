package block

import (
	"strconv"
)

type WaxedCopperBulb struct {
	Lit bool
	Powered bool
}

func (b WaxedCopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:waxed_copper_bulb", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WaxedCopperBulb) New(props BlockProperties) Block {
	return WaxedCopperBulb{
		Powered: props["powered"] != "false",
		Lit: props["lit"] != "false",
	}
}