package block

import (
	"strconv"
)

type WeepingVines struct {
	Age int
}

func (b WeepingVines) Encode() (string, BlockProperties) {
	return "minecraft:weeping_vines", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b WeepingVines) New(props BlockProperties) Block {
	return WeepingVines{
		Age: atoi(props["age"]),
	}
}