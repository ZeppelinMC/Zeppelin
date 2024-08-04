package block

import (
	"strconv"
)

type ZombieWallHead struct {
	Facing string
	Powered bool
}

func (b ZombieWallHead) Encode() (string, BlockProperties) {
	return "minecraft:zombie_wall_head", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b ZombieWallHead) New(props BlockProperties) Block {
	return ZombieWallHead{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}