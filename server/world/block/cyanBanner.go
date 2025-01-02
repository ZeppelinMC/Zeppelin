package block

import (
	"strconv"
)

type CyanBanner struct {
	Rotation int
}

func (b CyanBanner) Encode() (string, BlockProperties) {
	return "minecraft:cyan_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b CyanBanner) New(props BlockProperties) Block {
	return CyanBanner{
		Rotation: atoi(props["rotation"]),
	}
}