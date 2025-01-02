package block

import (
	"strconv"
)

type DragonWallHead struct {
	Facing string
	Powered bool
}

func (b DragonWallHead) Encode() (string, BlockProperties) {
	return "minecraft:dragon_wall_head", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b DragonWallHead) New(props BlockProperties) Block {
	return DragonWallHead{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}