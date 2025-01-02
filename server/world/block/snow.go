package block

import (
	"strconv"
)

type Snow struct {
	Layers int
}

func (b Snow) Encode() (string, BlockProperties) {
	return "minecraft:snow", BlockProperties{
		"layers": strconv.Itoa(b.Layers),
	}
}

func (b Snow) New(props BlockProperties) Block {
	return Snow{
		Layers: atoi(props["layers"]),
	}
}