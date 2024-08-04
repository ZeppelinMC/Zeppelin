package block

import (
	"strconv"
)

type WaterCauldron struct {
	Level int
}

func (b WaterCauldron) Encode() (string, BlockProperties) {
	return "minecraft:water_cauldron", BlockProperties{
		"level": strconv.Itoa(b.Level),
	}
}

func (b WaterCauldron) New(props BlockProperties) Block {
	return WaterCauldron{
		Level: atoi(props["level"]),
	}
}