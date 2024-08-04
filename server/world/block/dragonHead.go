package block

import (
	"strconv"
)

type DragonHead struct {
	Powered bool
	Rotation int
}

func (b DragonHead) Encode() (string, BlockProperties) {
	return "minecraft:dragon_head", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b DragonHead) New(props BlockProperties) Block {
	return DragonHead{
		Powered: props["powered"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}