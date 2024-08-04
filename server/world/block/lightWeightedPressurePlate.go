package block

import (
	"strconv"
)

type LightWeightedPressurePlate struct {
	Power int
}

func (b LightWeightedPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:light_weighted_pressure_plate", BlockProperties{
		"power": strconv.Itoa(b.Power),
	}
}

func (b LightWeightedPressurePlate) New(props BlockProperties) Block {
	return LightWeightedPressurePlate{
		Power: atoi(props["power"]),
	}
}