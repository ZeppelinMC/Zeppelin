package block

import (
	"strconv"
)

type WaxedOxidizedCopperBulb struct {
	Lit bool
	Powered bool
}

func (b WaxedOxidizedCopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:waxed_oxidized_copper_bulb", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WaxedOxidizedCopperBulb) New(props BlockProperties) Block {
	return WaxedOxidizedCopperBulb{
		Lit: props["lit"] != "false",
		Powered: props["powered"] != "false",
	}
}