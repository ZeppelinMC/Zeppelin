package block

import (
	"strconv"
)

type BirchPressurePlate struct {
	Powered bool
}

func (b BirchPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:birch_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b BirchPressurePlate) New(props BlockProperties) Block {
	return BirchPressurePlate{
		Powered: props["powered"] != "false",
	}
}