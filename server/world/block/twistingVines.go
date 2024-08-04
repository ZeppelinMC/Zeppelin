package block

import (
	"strconv"
)

type TwistingVines struct {
	Age int
}

func (b TwistingVines) Encode() (string, BlockProperties) {
	return "minecraft:twisting_vines", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b TwistingVines) New(props BlockProperties) Block {
	return TwistingVines{
		Age: atoi(props["age"]),
	}
}