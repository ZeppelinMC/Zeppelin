package block

import (
	"strconv"
)

type SugarCane struct {
	Age int
}

func (b SugarCane) Encode() (string, BlockProperties) {
	return "minecraft:sugar_cane", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b SugarCane) New(props BlockProperties) Block {
	return SugarCane{
		Age: atoi(props["age"]),
	}
}