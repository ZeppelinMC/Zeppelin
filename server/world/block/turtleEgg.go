package block

import (
	"strconv"
)

type TurtleEgg struct {
	Eggs int
	Hatch int
}

func (b TurtleEgg) Encode() (string, BlockProperties) {
	return "minecraft:turtle_egg", BlockProperties{
		"eggs": strconv.Itoa(b.Eggs),
		"hatch": strconv.Itoa(b.Hatch),
	}
}

func (b TurtleEgg) New(props BlockProperties) Block {
	return TurtleEgg{
		Eggs: atoi(props["eggs"]),
		Hatch: atoi(props["hatch"]),
	}
}