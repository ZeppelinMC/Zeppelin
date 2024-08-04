package block

import (
	"strconv"
)

type PinkPetals struct {
	Facing string
	FlowerAmount int
}

func (b PinkPetals) Encode() (string, BlockProperties) {
	return "minecraft:pink_petals", BlockProperties{
		"facing": b.Facing,
		"flower_amount": strconv.Itoa(b.FlowerAmount),
	}
}

func (b PinkPetals) New(props BlockProperties) Block {
	return PinkPetals{
		Facing: props["facing"],
		FlowerAmount: atoi(props["flower_amount"]),
	}
}