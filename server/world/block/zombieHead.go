package block

import (
	"strconv"
)

type ZombieHead struct {
	Powered bool
	Rotation int
}

func (b ZombieHead) Encode() (string, BlockProperties) {
	return "minecraft:zombie_head", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b ZombieHead) New(props BlockProperties) Block {
	return ZombieHead{
		Powered: props["powered"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}