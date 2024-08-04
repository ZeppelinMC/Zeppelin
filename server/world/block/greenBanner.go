package block

import (
	"strconv"
)

type GreenBanner struct {
	Rotation int
}

func (b GreenBanner) Encode() (string, BlockProperties) {
	return "minecraft:green_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b GreenBanner) New(props BlockProperties) Block {
	return GreenBanner{
		Rotation: atoi(props["rotation"]),
	}
}