package block

import (
	"strconv"
)

type LimeBanner struct {
	Rotation int
}

func (b LimeBanner) Encode() (string, BlockProperties) {
	return "minecraft:lime_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b LimeBanner) New(props BlockProperties) Block {
	return LimeBanner{
		Rotation: atoi(props["rotation"]),
	}
}