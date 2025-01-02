package block

import (
	"strconv"
)

type CherryPressurePlate struct {
	Powered bool
}

func (b CherryPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:cherry_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CherryPressurePlate) New(props BlockProperties) Block {
	return CherryPressurePlate{
		Powered: props["powered"] != "false",
	}
}