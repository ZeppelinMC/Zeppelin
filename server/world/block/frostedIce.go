package block

import (
	"strconv"
)

type FrostedIce struct {
	Age int
}

func (b FrostedIce) Encode() (string, BlockProperties) {
	return "minecraft:frosted_ice", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b FrostedIce) New(props BlockProperties) Block {
	return FrostedIce{
		Age: atoi(props["age"]),
	}
}