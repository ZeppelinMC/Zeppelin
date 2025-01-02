package block

import (
	"strconv"
)

type CreeperHead struct {
	Powered bool
	Rotation int
}

func (b CreeperHead) Encode() (string, BlockProperties) {
	return "minecraft:creeper_head", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CreeperHead) New(props BlockProperties) Block {
	return CreeperHead{
		Powered: props["powered"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}