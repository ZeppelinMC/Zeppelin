package block

import (
	"strconv"
)

type ChorusFlower struct {
	Age int
}

func (b ChorusFlower) Encode() (string, BlockProperties) {
	return "minecraft:chorus_flower", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b ChorusFlower) New(props BlockProperties) Block {
	return ChorusFlower{
		Age: atoi(props["age"]),
	}
}