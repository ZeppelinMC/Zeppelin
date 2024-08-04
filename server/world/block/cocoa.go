package block

import (
	"strconv"
)

type Cocoa struct {
	Age int
	Facing string
}

func (b Cocoa) Encode() (string, BlockProperties) {
	return "minecraft:cocoa", BlockProperties{
		"age": strconv.Itoa(b.Age),
		"facing": b.Facing,
	}
}

func (b Cocoa) New(props BlockProperties) Block {
	return Cocoa{
		Age: atoi(props["age"]),
		Facing: props["facing"],
	}
}