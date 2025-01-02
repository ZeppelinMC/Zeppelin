package block

import (
	"strconv"
)

type BlackBanner struct {
	Rotation int
}

func (b BlackBanner) Encode() (string, BlockProperties) {
	return "minecraft:black_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b BlackBanner) New(props BlockProperties) Block {
	return BlackBanner{
		Rotation: atoi(props["rotation"]),
	}
}