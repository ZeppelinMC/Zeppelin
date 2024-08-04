package block

import (
	"strconv"
)

type OrangeBanner struct {
	Rotation int
}

func (b OrangeBanner) Encode() (string, BlockProperties) {
	return "minecraft:orange_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b OrangeBanner) New(props BlockProperties) Block {
	return OrangeBanner{
		Rotation: atoi(props["rotation"]),
	}
}