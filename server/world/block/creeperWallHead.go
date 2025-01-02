package block

import (
	"strconv"
)

type CreeperWallHead struct {
	Facing string
	Powered bool
}

func (b CreeperWallHead) Encode() (string, BlockProperties) {
	return "minecraft:creeper_wall_head", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
	}
}

func (b CreeperWallHead) New(props BlockProperties) Block {
	return CreeperWallHead{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}