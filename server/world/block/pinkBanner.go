package block

import (
	"strconv"
)

type PinkBanner struct {
	Rotation int
}

func (b PinkBanner) Encode() (string, BlockProperties) {
	return "minecraft:pink_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b PinkBanner) New(props BlockProperties) Block {
	return PinkBanner{
		Rotation: atoi(props["rotation"]),
	}
}