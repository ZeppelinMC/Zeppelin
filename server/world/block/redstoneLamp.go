package block

import (
	"strconv"
)

type RedstoneLamp struct {
	Lit bool
}

func (b RedstoneLamp) Encode() (string, BlockProperties) {
	return "minecraft:redstone_lamp", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b RedstoneLamp) New(props BlockProperties) Block {
	return RedstoneLamp{
		Lit: props["lit"] != "false",
	}
}