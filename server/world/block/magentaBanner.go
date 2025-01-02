package block

import (
	"strconv"
)

type MagentaBanner struct {
	Rotation int
}

func (b MagentaBanner) Encode() (string, BlockProperties) {
	return "minecraft:magenta_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b MagentaBanner) New(props BlockProperties) Block {
	return MagentaBanner{
		Rotation: atoi(props["rotation"]),
	}
}