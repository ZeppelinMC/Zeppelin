package block

import (
	"strconv"
)

type PlayerHead struct {
	Powered bool
	Rotation int
}

func (b PlayerHead) Encode() (string, BlockProperties) {
	return "minecraft:player_head", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b PlayerHead) New(props BlockProperties) Block {
	return PlayerHead{
		Powered: props["powered"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}