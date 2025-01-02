package block

import (
	"strconv"
)

type Mycelium struct {
	Snowy bool
}

func (b Mycelium) Encode() (string, BlockProperties) {
	return "minecraft:mycelium", BlockProperties{
		"snowy": strconv.FormatBool(b.Snowy),
	}
}

func (b Mycelium) New(props BlockProperties) Block {
	return Mycelium{
		Snowy: props["snowy"] != "false",
	}
}