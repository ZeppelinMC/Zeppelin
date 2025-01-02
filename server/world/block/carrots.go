package block

import (
	"strconv"
)

type Carrots struct {
	Age int
}

func (b Carrots) Encode() (string, BlockProperties) {
	return "minecraft:carrots", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b Carrots) New(props BlockProperties) Block {
	return Carrots{
		Age: atoi(props["age"]),
	}
}