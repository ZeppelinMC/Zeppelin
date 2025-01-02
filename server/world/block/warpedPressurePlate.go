package block

import (
	"strconv"
)

type WarpedPressurePlate struct {
	Powered bool
}

func (b WarpedPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:warped_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WarpedPressurePlate) New(props BlockProperties) Block {
	return WarpedPressurePlate{
		Powered: props["powered"] != "false",
	}
}