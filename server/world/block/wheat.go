package block

import (
	"strconv"
)

type Wheat struct {
	Age int
}

func (b Wheat) Encode() (string, BlockProperties) {
	return "minecraft:wheat", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b Wheat) New(props BlockProperties) Block {
	return Wheat{
		Age: atoi(props["age"]),
	}
}