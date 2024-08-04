package block

import (
	"strconv"
)

type GrayBanner struct {
	Rotation int
}

func (b GrayBanner) Encode() (string, BlockProperties) {
	return "minecraft:gray_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b GrayBanner) New(props BlockProperties) Block {
	return GrayBanner{
		Rotation: atoi(props["rotation"]),
	}
}