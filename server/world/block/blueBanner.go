package block

import (
	"strconv"
)

type BlueBanner struct {
	Rotation int
}

func (b BlueBanner) Encode() (string, BlockProperties) {
	return "minecraft:blue_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b BlueBanner) New(props BlockProperties) Block {
	return BlueBanner{
		Rotation: atoi(props["rotation"]),
	}
}