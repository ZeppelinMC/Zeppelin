package block

import (
	"strconv"
)

type PowderSnowCauldron struct {
	Level int
}

func (b PowderSnowCauldron) Encode() (string, BlockProperties) {
	return "minecraft:powder_snow_cauldron", BlockProperties{
		"level": strconv.Itoa(b.Level),
	}
}

func (b PowderSnowCauldron) New(props BlockProperties) Block {
	return PowderSnowCauldron{
		Level: atoi(props["level"]),
	}
}