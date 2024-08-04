package block

import (
	"strconv"
)

type StonePressurePlate struct {
	Powered bool
}

func (b StonePressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:stone_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b StonePressurePlate) New(props BlockProperties) Block {
	return StonePressurePlate{
		Powered: props["powered"] != "false",
	}
}