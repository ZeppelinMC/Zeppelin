package block

import (
	"strconv"
)

type HeavyWeightedPressurePlate struct {
	Power int
}

func (b HeavyWeightedPressurePlate) Encode() (string, BlockProperties) {
	return "minecraft:heavy_weighted_pressure_plate", BlockProperties{
		"power": strconv.Itoa(b.Power),
	}
}

func (b HeavyWeightedPressurePlate) New(props BlockProperties) Block {
	return HeavyWeightedPressurePlate{
		Power: atoi(props["power"]),
	}
}