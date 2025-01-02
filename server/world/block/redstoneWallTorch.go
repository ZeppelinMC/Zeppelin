package block

import (
	"strconv"
)

type RedstoneWallTorch struct {
	Facing string
	Lit bool
}

func (b RedstoneWallTorch) Encode() (string, BlockProperties) {
	return "minecraft:redstone_wall_torch", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"facing": b.Facing,
	}
}

func (b RedstoneWallTorch) New(props BlockProperties) Block {
	return RedstoneWallTorch{
		Facing: props["facing"],
		Lit: props["lit"] != "false",
	}
}