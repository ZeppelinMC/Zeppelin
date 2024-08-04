package block

import (
	"strconv"
)

type LightBlueBanner struct {
	Rotation int
}

func (b LightBlueBanner) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b LightBlueBanner) New(props BlockProperties) Block {
	return LightBlueBanner{
		Rotation: atoi(props["rotation"]),
	}
}