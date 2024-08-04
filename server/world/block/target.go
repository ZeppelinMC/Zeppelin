package block

import (
	"strconv"
)

type Target struct {
	Power int
}

func (b Target) Encode() (string, BlockProperties) {
	return "minecraft:target", BlockProperties{
		"power": strconv.Itoa(b.Power),
	}
}

func (b Target) New(props BlockProperties) Block {
	return Target{
		Power: atoi(props["power"]),
	}
}