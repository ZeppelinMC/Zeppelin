package block

import (
	"strconv"
)

type BeeNest struct {
	Facing string
	HoneyLevel int
}

func (b BeeNest) Encode() (string, BlockProperties) {
	return "minecraft:bee_nest", BlockProperties{
		"facing": b.Facing,
		"honey_level": strconv.Itoa(b.HoneyLevel),
	}
}

func (b BeeNest) New(props BlockProperties) Block {
	return BeeNest{
		Facing: props["facing"],
		HoneyLevel: atoi(props["honey_level"]),
	}
}