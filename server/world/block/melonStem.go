package block

import (
	"strconv"
)

type MelonStem struct {
	Age int
}

func (b MelonStem) Encode() (string, BlockProperties) {
	return "minecraft:melon_stem", BlockProperties{
		"age": strconv.Itoa(b.Age),
	}
}

func (b MelonStem) New(props BlockProperties) Block {
	return MelonStem{
		Age: atoi(props["age"]),
	}
}