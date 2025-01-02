package block

import (
	"strconv"
)

type CaveVinesPlant struct {
	Berries bool
}

func (b CaveVinesPlant) Encode() (string, BlockProperties) {
	return "minecraft:cave_vines_plant", BlockProperties{
		"berries": strconv.FormatBool(b.Berries),
	}
}

func (b CaveVinesPlant) New(props BlockProperties) Block {
	return CaveVinesPlant{
		Berries: props["berries"] != "false",
	}
}