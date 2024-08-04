package block

import (
	"strconv"
)

type PiglinHead struct {
	Powered bool
	Rotation int
}

func (b PiglinHead) Encode() (string, BlockProperties) {
	return "minecraft:piglin_head", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b PiglinHead) New(props BlockProperties) Block {
	return PiglinHead{
		Powered: props["powered"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}