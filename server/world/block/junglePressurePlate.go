package block

import (
	"strconv"
)

type JunglePressurePlate struct {
	Powered bool
}

func (b JunglePressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:jungle_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b JunglePressurePlate) New(props BlockProperties) Block {
	return JunglePressurePlate{
		Powered: props["powered"] != "false",
	}
}