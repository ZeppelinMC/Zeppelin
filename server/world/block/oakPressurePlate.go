package block

import (
	"strconv"
)

type OakPressurePlate struct {
	Powered bool
}

func (b OakPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:oak_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b OakPressurePlate) New(props BlockProperties) Block {
	return OakPressurePlate{
		Powered: props["powered"] != "false",
	}
}