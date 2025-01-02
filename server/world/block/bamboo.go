package block

import (
	"strconv"
)

type Bamboo struct {
	Age int
	Leaves string
	Stage int
}

func (b Bamboo) Encode() (string, BlockProperties) {
	return "minecraft:bamboo", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
		"age": strconv.Itoa(b.Age),
		"leaves": b.Leaves,
	}
}

func (b Bamboo) New(props BlockProperties) Block {
	return Bamboo{
		Age: atoi(props["age"]),
		Leaves: props["leaves"],
		Stage: atoi(props["stage"]),
	}
}