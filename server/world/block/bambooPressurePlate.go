package block

import (
	"strconv"
)

type BambooPressurePlate struct {
	Powered bool
}

func (b BambooPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b BambooPressurePlate) New(props BlockProperties) Block {
	return BambooPressurePlate{
		Powered: props["powered"] != "false",
	}
}