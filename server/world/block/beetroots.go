package block

import (
	"strconv"
)

type Beetroots struct {
	Age int
}

func (b Beetroots) Encode() (string, BlockProperties) {
	return "minecraft:beetroots", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b Beetroots) New(props BlockProperties) Block {
	return Beetroots{
		Age: atoi(props["age"]),
	}
}