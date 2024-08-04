package block

import (
	"strconv"
)

type NetherWart struct {
	Age int
}

func (b NetherWart) Encode() (string, BlockProperties) {
	return "minecraft:nether_wart", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b NetherWart) New(props BlockProperties) Block {
	return NetherWart{
		Age: atoi(props["age"]),
	}
}