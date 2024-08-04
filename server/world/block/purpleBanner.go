package block

import (
	"strconv"
)

type PurpleBanner struct {
	Rotation int
}

func (b PurpleBanner) Encode() (string, BlockProperties) {
	return "minecraft:purple_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b PurpleBanner) New(props BlockProperties) Block {
	return PurpleBanner{
		Rotation: atoi(props["rotation"]),
	}
}