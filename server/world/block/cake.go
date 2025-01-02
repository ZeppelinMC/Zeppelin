package block

import (
	"strconv"
)

type Cake struct {
	Bites int
}

func (b Cake) Encode() (string, BlockProperties) {
	return "minecraft:cake", BlockProperties{
		"bites": strconv.Itoa(b.Bites),
	}
}

func (b Cake) New(props BlockProperties) Block {
	return Cake{
		Bites: atoi(props["bites"]),
	}
}