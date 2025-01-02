package block

import (
	"strconv"
)

type CopperBulb struct {
	Lit bool
	Powered bool
}

func (b CopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:copper_bulb", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b CopperBulb) New(props BlockProperties) Block {
	return CopperBulb{
		Lit: props["lit"] != "false",
		Powered: props["powered"] != "false",
	}
}