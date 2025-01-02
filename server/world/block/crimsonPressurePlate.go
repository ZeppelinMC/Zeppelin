package block

import (
	"strconv"
)

type CrimsonPressurePlate struct {
	Powered bool
}

func (b CrimsonPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:crimson_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CrimsonPressurePlate) New(props BlockProperties) Block {
	return CrimsonPressurePlate{
		Powered: props["powered"] != "false",
	}
}