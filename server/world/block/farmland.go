package block

import (
	"strconv"
)

type Farmland struct {
	Moisture int
}

func (b Farmland) Encode() (string, BlockProperties) {
	return "minecraft:farmland", BlockProperties{
		"moisture": strconv.Itoa(b.Moisture),
	}
}

func (b Farmland) New(props BlockProperties) Block {
	return Farmland{
		Moisture: atoi(props["moisture"]),
	}
}