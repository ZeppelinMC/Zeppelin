package block

import (
	"strconv"
)

type SweetBerryBush struct {
	Age int
}

func (b SweetBerryBush) Encode() (string, BlockProperties) {
	return "minecraft:sweet_berry_bush", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b SweetBerryBush) New(props BlockProperties) Block {
	return SweetBerryBush{
		Age: atoi(props["age"]),
	}
}