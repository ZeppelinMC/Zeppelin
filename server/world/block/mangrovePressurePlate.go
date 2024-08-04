package block

import (
	"strconv"
)

type MangrovePressurePlate struct {
	Powered bool
}

func (b MangrovePressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b MangrovePressurePlate) New(props BlockProperties) Block {
	return MangrovePressurePlate{
		Powered: props["powered"] != "false",
	}
}