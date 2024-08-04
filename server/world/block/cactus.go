package block

import (
	"strconv"
)

type Cactus struct {
	Age int
}

func (b Cactus) Encode() (string, BlockProperties) {
	return "minecraft:cactus", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b Cactus) New(props BlockProperties) Block {
	return Cactus{
		Age: atoi(props["age"]),
	}
}