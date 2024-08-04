package block

import (
	"strconv"
)

type RedBanner struct {
	Rotation int
}

func (b RedBanner) Encode() (string, BlockProperties) {
	return "minecraft:red_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b RedBanner) New(props BlockProperties) Block {
	return RedBanner{
		Rotation: atoi(props["rotation"]),
	}
}