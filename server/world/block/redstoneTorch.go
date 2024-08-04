package block

import (
	"strconv"
)

type RedstoneTorch struct {
	Lit bool
}

func (b RedstoneTorch) Encode() (string, BlockProperties) {
	return "minecraft:redstone_torch", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b RedstoneTorch) New(props BlockProperties) Block {
	return RedstoneTorch{
		Lit: props["lit"] != "false",
	}
}