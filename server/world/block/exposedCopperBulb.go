package block

import (
	"strconv"
)

type ExposedCopperBulb struct {
	Powered bool
	Lit bool
}

func (b ExposedCopperBulb) Encode() (string, BlockProperties) {
	return "minecraft:exposed_copper_bulb", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b ExposedCopperBulb) New(props BlockProperties) Block {
	return ExposedCopperBulb{
		Powered: props["powered"] != "false",
		Lit: props["lit"] != "false",
	}
}