package block

import (
	"strconv"
)

type PlayerWallHead struct {
	Powered bool
	Facing string
}

func (b PlayerWallHead) Encode() (string, BlockProperties) {
	return "minecraft:player_wall_head", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b PlayerWallHead) New(props BlockProperties) Block {
	return PlayerWallHead{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}