package block

import (
	"strconv"
)

type PiglinWallHead struct {
	Facing string
	Powered bool
}

func (b PiglinWallHead) Encode() (string, BlockProperties) {
	return "minecraft:piglin_wall_head", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b PiglinWallHead) New(props BlockProperties) Block {
	return PiglinWallHead{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
	}
}