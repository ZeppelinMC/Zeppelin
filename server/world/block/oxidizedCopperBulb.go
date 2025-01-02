package block

import (
	"strconv"
)

type OxidizedCopperBulb struct {
	Lit bool
	Powered bool
}

func (b OxidizedCopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:oxidized_copper_bulb", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b OxidizedCopperBulb) New(props BlockProperties) Block {
	return OxidizedCopperBulb{
		Powered: props["powered"] != "false",
		Lit: props["lit"] != "false",
	}
}