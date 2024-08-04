package block

import (
	"strconv"
)

type Podzol struct {
	Snowy bool
}

func (b Podzol) Encode() (string, BlockProperties) {
	return "minecraft:podzol", BlockProperties{
		"snowy": strconv.FormatBool(b.Snowy),
	}
}

func (b Podzol) New(props BlockProperties) Block {
	return Podzol{
		Snowy: props["snowy"] != "false",
	}
}