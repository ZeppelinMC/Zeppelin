package block

import (
	"strconv"
)

type WhiteBanner struct {
	Rotation int
}

func (b WhiteBanner) Encode() (string, BlockProperties) {
	return "minecraft:white_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b WhiteBanner) New(props BlockProperties) Block {
	return WhiteBanner{
		Rotation: atoi(props["rotation"]),
	}
}