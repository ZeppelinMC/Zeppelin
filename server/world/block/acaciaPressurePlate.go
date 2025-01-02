package block

import (
	"strconv"
)

type AcaciaPressurePlate struct {
	Powered bool
}

func (b AcaciaPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:acacia_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b AcaciaPressurePlate) New(props BlockProperties) Block {
	return AcaciaPressurePlate{
		Powered: props["powered"] != "false",
	}
}