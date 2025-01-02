package block

import (
	"strconv"
)

type PumpkinStem struct {
	Age int
}

func (b PumpkinStem) Encode() (string, BlockProperties) {
	return "minecraft:pumpkin_stem", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b PumpkinStem) New(props BlockProperties) Block {
	return PumpkinStem{
		Age: atoi(props["age"]),
	}
}