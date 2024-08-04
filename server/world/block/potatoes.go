package block

import (
	"strconv"
)

type Potatoes struct {
	Age int
}

func (b Potatoes) Encode() (string, BlockProperties) {
	return "minecraft:potatoes", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b Potatoes) New(props BlockProperties) Block {
	return Potatoes{
		Age: atoi(props["age"]),
	}
}