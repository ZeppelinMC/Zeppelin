package block

import (
	"strconv"
)

type DeepslateRedstoneOre struct {
	Lit bool
}

func (b DeepslateRedstoneOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_redstone_ore", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b DeepslateRedstoneOre) New(props BlockProperties) Block {
	return DeepslateRedstoneOre{
		Lit: props["lit"] != "false",
	}
}