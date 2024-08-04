package block

import (
	"strconv"
)

type YellowBanner struct {
	Rotation int
}

func (b YellowBanner) Encode() (string, BlockProperties) {
	return "minecraft:yellow_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b YellowBanner) New(props BlockProperties) Block {
	return YellowBanner{
		Rotation: atoi(props["rotation"]),
	}
}