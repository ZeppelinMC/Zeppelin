package block

import (
	"strconv"
)

type RedstoneOre struct {
	Lit bool
}

func (b RedstoneOre) Encode() (string, BlockProperties) {
	return "minecraft:redstone_ore", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b RedstoneOre) New(props BlockProperties) Block {
	return RedstoneOre{
		Lit: props["lit"] != "false",
	}
}