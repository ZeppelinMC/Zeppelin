package block

import (
	"strconv"
)

type PolishedBlackstonePressurePlate struct {
	Powered bool
}

func (b PolishedBlackstonePressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone_pressure_plate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b PolishedBlackstonePressurePlate) New(props BlockProperties) Block {
	return PolishedBlackstonePressurePlate{
		Powered: props["powered"] != "false",
	}
}