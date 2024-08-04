package block

import (
	"strconv"
)

type DarkOakPressurePlate struct {
	Powered bool
}

func (b DarkOakPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b DarkOakPressurePlate) New(props BlockProperties) Block {
	return DarkOakPressurePlate{
		Powered: props["powered"] != "false",
	}
}