package block

import (
	"strconv"
)

type SprucePressurePlate struct {
	Powered bool
}

func (b SprucePressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:spruce_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b SprucePressurePlate) New(props BlockProperties) Block {
	return SprucePressurePlate{
		Powered: props["powered"] != "false",
	}
}