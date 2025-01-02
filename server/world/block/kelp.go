package block

import (
	"strconv"
)

type Kelp struct {
	Age int
}

func (b Kelp) Encode() (string, BlockProperties) {
	return "minecraft:kelp", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b Kelp) New(props BlockProperties) Block {
	return Kelp{
		Age: atoi(props["age"]),
	}
}