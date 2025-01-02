package block

import (
	"strconv"
)

type CaveVines struct {
	Age int
	Berries bool
}

func (b CaveVines) Encode() (string, BlockProperties) {
	return "minecraft:cave_vines", BlockProperties{
		"berries": strconv.FormatBool(b.Berries),
		"age": strconv.Itoa(b.Age),
	}
}

func (b CaveVines) New(props BlockProperties) Block {
	return CaveVines{
		Age: atoi(props["age"]),
		Berries: props["berries"] != "false",
	}
}